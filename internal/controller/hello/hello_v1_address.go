package hello

import (
	"ccarbon-coin/internal/service"
	"context"

	"ccarbon-coin/api/hello/v1"
)

func (c *ControllerV1) Address(ctx context.Context, req *v1.AddressReq) (res *v1.AddressRes, err error) {
	address, err := service.Coin().GenAddress(ctx, req.Flag)
	if err != nil {
		return
	}
	res = &v1.AddressRes{
		EthAddress: address.GetVar("eth_address").String(),
		TrxAddress: address.GetVar("trx_address").String(),
		SolAddress: address.GetVar("sol_address").String(),
	}
	return
}
