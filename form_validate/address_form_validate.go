package form_validate

type AddressForm struct {
	Id           int64      `form:"id"`
	AssetId      int64      `form:"asset_id"`
	DeviceId     string     `form:"device_id"`
	WalletUuid   string     `form:"wallet_uuid"`
	WalletName   string     `form:"wallet_name"`
	Address      string     `form:"address"`
	ContractAddr string     `form:"contract_addr"`
	PrivateKey   string     `form:"private_key"`
	Balance      float64    `form:"balance"`
	IsCreate     int        `form:"_create"`
}