package model

type CoinMsg struct {
	MsgType  int    `json:"msg_type"`
	CoinType int    `json:"cointype"`
	Tx       string `json:"tx"`
	From     string `json:"from"`
	To       string `json:"to"`
	Amount   string `json:"amount"`
	Decimal  int    `json:"decimal"`
	Symbol   int    `json:"symbol"`
	Contract string `json:"contract"`
	Height   string `json:"height"`
	User     string `json:"user"`
}

type CoinOutReq struct {
	CoinType int    `json:"coin_type"`
	Amount   string `json:"amount"`
	Address  string `json:"address"`
	User     string `json:"user"`
}
