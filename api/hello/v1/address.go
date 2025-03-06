package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddressReq struct {
	g.Meta `path:"/address" tags:"Hello" method:"post" summary:"Gen address"`
	Flag   string `p:"flag" v:"required"`
}
type AddressRes struct {
	EthAddress string `json:"eth_address"`
	TrxAddress string `json:"trx_address"`
	SolAddress string `json:"sol_address"`
}
