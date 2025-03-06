// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import (
	"context"

	"ccarbon-coin/api/hello/v1"
)

type IHelloV1 interface {
	Address(ctx context.Context, req *v1.AddressReq) (res *v1.AddressRes, err error)
	Withdraw(ctx context.Context, req *v1.WithdrawReq) (res *v1.WithdrawRes, err error)
}
