package token

type TokenSoarchRep struct {
	Id           int64 `json:"id"`
	AssetId      int64 `json:"asset_id"`
	TokenName    string `json:"token_name"`
	Icon         string `json:"icon"`
	TokenSymbol  string `json:"token_symbol"`
	ContractAddr string `json:"contract_addr"`
	Decimal      int64 `json:"decimal"`
}
