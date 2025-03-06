package coin

import (
	"ccarbon-coin/internal/model"
	"context"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net/http"
	"time"
)

func startWss() {
	ctx := gctx.New()
	retryInterval := time.Second * 5

	for {
		token, err := genToken(ctx)
		if err != nil {
			g.Log().Error(ctx, "Failed to generate token:", err)
			return
		}

		client := gclient.NewWebSocket()
		client.HandshakeTimeout = time.Second * 5
		header := http.Header{}
		header.Set("token", token)

		conn, resp, err := client.Dial(cfg.Wss+"?client_id="+cfg.ClientId, header)
		if err != nil {
			if resp != nil {
				g.Log().Error(ctx, "Handshake failed with status:", resp.Status)
				body, _ := io.ReadAll(resp.Body)
				g.Log().Debug(ctx, "Response body:", string(body))
			}
			g.Log().Error(ctx, "Connection error:", err)
			time.Sleep(retryInterval)
			g.Log().Debug(ctx, "Try to reconnect...")
			continue
		}

		g.Log().Debug(ctx, "Coin connection successful")
		defer conn.Close()

		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				g.Log().Error(ctx, "ReadMessage error:", err)
				break
			}
			txs := g.ArrayStr{}
			txs, _ = msgHandler(string(data))
			if len(txs) > 0 {
				_ = conn.WriteJSON(g.Map{"tx": txs, "type": "confirm"})
			}
		}
		g.Log().Info(ctx, "Connection lost. Reconnecting in", retryInterval)
		time.Sleep(retryInterval)
	}
}

func msgHandler(dataStr string) (txs g.ArrayStr, err error) {
	ctx := gctx.New()
	g.Log().Info(ctx, "收到消息:", dataStr)
	d := gmap.NewStrAnyMap()
	err = gconv.Scan(dataStr, &d)
	if err != nil {
		return
	}
	data := d.GetVar("data").Array()
	for _, v := range data {
		one := model.CoinMsg{}
		err = gconv.Scan(v, &one)
		if err != nil {
			g.Log().Error(ctx, "Message parsing error:", err)
			continue
		}
		one.MsgType = d.GetVar("type").Int()
		err = cb(ctx, &one)
		if err != nil {
			g.Log().Error(ctx, "Callback error:", err)
			continue
		}
		txs = append(txs, one.Tx)
	}
	return
}

func cb(ctx context.Context, data *model.CoinMsg) (err error) {
	url := g.Cfg().MustGet(ctx, "forward.url").String()
	r, err := g.Client().
		SetHeader("Content-Type", "application/json").
		SetHeader("CoinKey", g.Cfg().MustGet(ctx, "forward.key").String()).
		Post(ctx, url, data)
	if g.Cfg().MustGet(ctx, "app.test").Bool() {
		r.RawDump()
	}
	if err != nil {
		return
	}
	defer r.Close()
	if r.StatusCode != 200 {
		err = gerror.Newf("Forward failed, status code: %d", r.StatusCode)
		return
	}
	resBody := r.ReadAllString()
	if resBody != "success" {
		err = gerror.New("Forward failed, response body: " + resBody)
	}
	return
}
