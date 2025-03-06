package coin

import (
	"ccarbon-coin/internal/model"
	"ccarbon-coin/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type (
	sCoin struct{}
)

var cfg struct {
	Domain   string
	Wss      string
	ClientId string
}

func init() {
	service.RegisterCoin(New())
	err := g.Cfg().MustGet(gctx.GetInitCtx(), "coin").Struct(&cfg)
	if err != nil {
		return
	}
	go startWss()
}

func New() service.ICoin {
	return &sCoin{}
}

func (s *sCoin) GenAddress(ctx context.Context, falg string) (res gmap.StrAnyMap, err error) {
	//1    ERC20 ETH
	//56  BEP20 BSC
	//137 MATIC POLYGON
	//195 TRC20 TRON
	//200 Solana Solana
	path := fmt.Sprintf("/user/generate?user=%s", falg)
	return senReq(ctx, path, nil)
}

func (s *sCoin) Withdraw(ctx context.Context, data model.CoinOutReq) (res gmap.StrAnyMap, err error) {
	p := gmap.NewStrAnyMapFrom(gconv.Map(data))
	return senReq(ctx, "/user/withdraw", p)
}

func genToken(ctx context.Context) (tokenStr string, err error) {
	cache, err := g.Redis().Get(ctx, "coin-token")
	if err != nil {
		return
	}
	if cache.String() != "" {
		tokenStr = cache.String()
		return
	}

	token := jwt.New(jwt.SigningMethodRS256)
	privatePath := g.Cfg().MustGet(ctx, "app.storage").String() + "/cert/coin_private.pem"
	privateKeyData := gfile.GetBytes(privatePath)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyData)
	if err != nil {
		return
	}
	tokenStr, err = token.SignedString(privateKey)
	if err != nil {
		return
	}
	_ = g.Redis().SetEX(ctx, "coin-token", tokenStr, int64(gtime.D*365/time.Second))
	return
}

func senReq(ctx context.Context, path string, data *gmap.StrAnyMap) (res gmap.StrAnyMap, err error) {
	token, err := genToken(ctx)
	if err != nil {
		return
	}
	if strings.Contains(path, "?") {
		path += "&client_id=" + cfg.ClientId
	} else {
		path += "?client_id=" + cfg.ClientId
	}
	url := cfg.Domain + path

	jsonStr, err := gjson.Encode(data)
	if err != nil {
		return
	}
	client := gclient.New()
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("token", token)
	r, err := client.Post(ctx, url, jsonStr)
	if g.Cfg().MustGet(ctx, "app.test").Bool() {
		r.RawDump()
	}
	if err != nil {
		g.Log().Error(ctx, "Coin 请求失败 #01", err)
		err = gerror.New("Coin 请求失败 #01")
		return
	}
	defer r.Close()

	resBody := r.ReadAllString()
	ret := gmap.NewStrAnyMap()
	err = gjson.DecodeTo(resBody, &ret)
	if ret.GetVar("code").Int() != 200 {
		g.Log().Error(ctx, "Coin 请求失败 #02", err)
		err = gerror.New("Coin 请求失败 #02")
		return
	}
	err = ret.GetVar("data").Struct(&res)
	if err != nil {
		return
	}
	return
}
