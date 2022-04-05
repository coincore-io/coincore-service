package wallet


type WalletBalanceResponse struct {
	Id        int64   `json:"id"`
	DeviceId   string `json:"device_id"`
	WalletUuid string `json:"wallet_uuid"`
	WalletName string `json:"wallet_name"`
	Balance   float64 `json:"balance"`
	Icon      string  `json:"icon"`
	Name      string  `json:"name"`
	ChainName string  `json:"chain_name"`
	Address   string  `json:"address"`
	ContractAddr string `json:"contract_addr"`
	UsdtPrice float64 `json:"usdt_price"`
	CnyPrice  float64 `json:"cny_price"`
}

type GasPriceResponse struct {
	Index    int64 `json:"index"`
	GasPrice int64 `json:"gas_price"`
}

type WalletAssetRep struct {
	WalletName string `json:"wallet_name"`
	WalletBalance []*WalletBalanceResponse `json:"wallet_balance"`
}

type AddressStatData struct {
	Amount  float64 `json:"amount"`
	DateTime    string `json:"date_time"`
	Time        string `json:"time"`
}


type TransRecordDateRep struct {
	BlockNumber 	string `json:"block_number"`
	DateTine		string `json:"date_tine"`
	AssetName       string `json:"asset_name"`
	Hash 			string `json:"hash"`
	From 			string `json:"from"`
	To 				string `json:"to"`
	Value 			string `json:"value"`
	ContractAddress string `json:"contract_address"`
	GasUsed 		string `json:"gas_used"`
	GasPrice        string `json:"gas_price"`
	IsError         string `json:"is_error"`
	TxreceiptStatus string `json:"txreceipt_status"`
	TxInOut         string `json:"tx_in_out"`
	Unit            int64  `json:"unit"`
}