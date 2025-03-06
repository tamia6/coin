// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"ccarbon-coin/internal/model"
	"context"

	"github.com/gogf/gf/v2/container/gmap"
)

type (
	ICoin interface {
		GenAddress(ctx context.Context, falg string) (res gmap.StrAnyMap, err error)
		Withdraw(ctx context.Context, data model.CoinOutReq) (res gmap.StrAnyMap, err error)
	}
)

var (
	localCoin ICoin
)

func Coin() ICoin {
	if localCoin == nil {
		panic("implement not found for interface ICoin, forgot register?")
	}
	return localCoin
}

func RegisterCoin(i ICoin) {
	localCoin = i
}
