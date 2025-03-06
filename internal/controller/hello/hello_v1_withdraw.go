package hello

import (
	"ccarbon-coin/internal/model"
	"ccarbon-coin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"ccarbon-coin/api/hello/v1"
)

func (c *ControllerV1) Withdraw(ctx context.Context, req *v1.WithdrawReq) (res *v1.WithdrawRes, err error) {
	g.Log().Info(ctx, "WithdrawReq:", req)
	_, err = service.Coin().Withdraw(ctx, model.CoinOutReq{
		CoinType: req.CoinType,
		Amount:   gconv.String(req.Amount),
		Address:  req.Address,
		User:     req.Flag,
	})
	if err != nil {
		return
	}
	return
}
