package form_validate

type TokenConfigForm struct {
	Id           int64      `form:"id"`
	AssetId      int64      `form:"asset_id"`
	TokenName    string     `form:"token_name"`
	Icon         string     `form:"icon"`
	TokenSymbol  string     `form:"token_symbol"`
	ContractAddr string     `form:"contract_addr"`
	Decimal      int64      `form:"decimal"`
	IsCreate     int        `form:"_create"`
}