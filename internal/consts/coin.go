package consts

type CoinTypeEnum int

const (
	//1    ERC20 ETH
	//56  BEP20 BSC
	//137 MATIC POLYGON
	//195 TRC20 TRON
	//200 Solana Solana

	ETH  CoinTypeEnum = 1
	TRON CoinTypeEnum = 195
)
