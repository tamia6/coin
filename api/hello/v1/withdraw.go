package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WithdrawReq struct {
	g.Meta   `path:"/withdraw" tags:"Hello" method:"post" summary:"withdraw"`
	Address  string  `p:"address" v:"required"`
	Amount   float64 `p:"amount" v:"required"`
	CoinType int     `p:"type" v:"required"`
	Flag     string  `p:"flag" v:"required"`
}
type WithdrawRes struct {
	Status bool `json:"status"`
}
