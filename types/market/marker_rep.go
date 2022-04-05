package market


type MarketPriceRep struct {
	Id        int64   `json:"id"`
	ChainName string  `json:"chain_name"`
	Name      string  `json:"name"`
	Icon      string  `json:"icon"`
	UsdPrice  float64 `json:"usd_price"`
	CnyPrice  float64 `json:"cny_price"`
	Rate      float64 `json:"rate"`
}
