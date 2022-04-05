package api_v1

import (
	"coinwallet/models"
	"coinwallet/types"
	"coinwallet/types/wallet"
	"coinwallet/wallet/eth"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"strconv"
	"time"
)

type WalletController struct {
	beego.Controller
}

type Int struct {
	i *big.Int
}

var ethClient_ *eth.EthClient

const (
	EtherScanBaseUrl = "https://api.etherscan.io/api"
	EthNode = "https://mainnet.infura.io/v3/b48b6387e66d4f3497245873747f6e4d"
	ApiKey = "5W7VEUM5UVVMK4CZ4ZSRSW1EUGX17SKI1Z"
	Startblock = "0"
	Endblock = "999999999"
	GasLimit = 90000
	maxBitLen = 255
)


func init()  {
	var err error
	ethClient_, err = eth.NewEthClients(EthNode)
	if err != nil {
		panic(err)
	}
}

// SubmitWalletInfo @Title SubmitWalletInfo
// @Description 提交钱包信息 SubmitWalletInfo
// @Success 200 status bool, data interface{}, msg string
// @router /submit_wallet_info [post]
func (this *WalletController) SubmitWalletInfo() {
	var sb_wallet_req wallet.SubmitWalletInfoReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &sb_wallet_req); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := sb_wallet_req.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	err := models.SubmitWalletInfo(sb_wallet_req)
	if err != nil {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "wallet success")
	this.ServeJSON()
	return
}


// BatchSubmitWallet @Title BatchSubmitWallet
// @Description 批量提交钱包信息 BatchSubmitWallet
// @Success 200 status bool, data interface{}, msg string
// @router /batch_submit_wallet [post]
func (this *WalletController) BatchSubmitWallet() {
	var batch_wallet wallet.BatchSubmitWalletReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &batch_wallet); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	fmt.Println(batch_wallet)
	if code, err := batch_wallet.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	for _, v := range batch_wallet.BatchWallet {
		err := models.SubmitWalletInfo(v)
		if err != nil {
			fmt.Println(err)
		}
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "wallet success")
	this.ServeJSON()
	return
}

// DeleteWallet @Title DeleteWallet
// @Description 删除钱包 DeleteWallet
// @Success 200 status bool, data interface{}, msg string
// @router /delete_wallet [post]
func (this *WalletController) DeleteWallet() {
	var delete_wt wallet.DeleteWalletReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &delete_wt); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := delete_wt.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	ok := models.DeleteWallet(delete_wt.DeviceId, delete_wt.WalletUuid)
	if ok == false {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "wallet success")
	this.ServeJSON()
	return
}

// GetAddressBalance @Title GetAddressBalance
// @Description 根据地址获取余额 GetAddressBalance
// @Success 200 status bool, data interface{}, msg string
// @router /get_address_balance [post]
func (this *WalletController) GetAddressBalance() {
	var address_balance wallet.AddressBalanceReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &address_balance); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := address_balance.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	m_asst := models.MarketAsset{}
	_ = orm.NewOrm().QueryTable(models.MarketAsset{}).Filter("name", address_balance.AssetName).One(&m_asst)
	if address_balance.ContractAddr != "" {
		balance_erc20, err := ethClient_.Erc20BalanceOf(address_balance.ContractAddr, address_balance.Address,nil)
		if err != nil {
			this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
			this.ServeJSON()
			return
		}
		unit, err := ethClient_.Erc20Decimals(address_balance.ContractAddr)
		if err != nil {
			this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
			this.ServeJSON()
			return
		}
		balance_decimal := decimal.NewFromBigInt(balance_erc20, 0).DivRound(decimal.New(10, 0).Pow(decimal.New(int64(unit), 0)), int32(unit))
		balance, _ := balance_decimal.Float64()
		models.UpdateBalanceByAddress(
			address_balance.AssetName,
			address_balance.ChainName,
			address_balance.DeviceId,
			address_balance.WalletUuid,
			address_balance.ContractAddr,
			balance,
		)
		if err != nil {
			this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
			this.ServeJSON()
			return
		}
		if err != nil {
			this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
			this.ServeJSON()
			return
		}
		balance_h := math.Trunc(balance*1e4+0.5) * 1e-4
		var usdt_price,  cny_price float64
		market_p := models.GetMarketByAssetId(m_asst.Id)
		if market_p != nil {
			usdt_price = balance_h * market_p.UsdPrice
			cny_price = balance_h * market_p.CnyPrice
		} else {
			usdt_price = 0
			cny_price = 0
		}

		wallet_stat_db_lst, _ := models.GetWalletStatWid(1)
		var wallet_stat_list []wallet.AddressStatData
		for _, wallet_stat := range wallet_stat_db_lst {
			w_sdata := wallet.AddressStatData{
				Amount: wallet_stat.Amount,
				DateTime:  wallet_stat.CreatedAt.Format("2006-01-02 15:04:05"),
				Time: wallet_stat.CreatedAt.Format("15:04"),
			}
			wallet_stat_list = append(wallet_stat_list, w_sdata)
		}
		data := map[string]interface{}{
			"balance": balance_h,
			"icon": m_asst.Icon,
			"name": m_asst.Name,
			"adddress": address_balance.Address,
			"chain_name": m_asst.ChainName,
			"usdt_price": usdt_price,
			"cny_price": cny_price ,
			"unit": unit,
			"data_stat": wallet_stat_list,
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "wallet success")
		this.ServeJSON()
		return
	} else {
		balance_big_start, err := ethClient_.GetEthBalance(address_balance.Address)
		if err != nil {
			this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
			this.ServeJSON()
			return
		}
		balance_decimal := decimal.NewFromBigInt(balance_big_start, 0).DivRound(decimal.New(10, 0).Pow(decimal.New(18, 0)), 18)
		balance, _ := balance_decimal.Float64()
		models.UpdateBalanceByAddress(
			address_balance.AssetName,
			address_balance.ChainName,
			address_balance.DeviceId,
			address_balance.WalletUuid,
			address_balance.ContractAddr,
			balance,
		)
		var usdt_price,  cny_price float64
		balance_h := math.Trunc(balance*1e4+0.5) * 1e-4
		market_p := models.GetMarketByAssetId(m_asst.Id)
		if market_p != nil {
			usdt_price = balance_h * market_p.UsdPrice
			cny_price = balance_h * market_p.CnyPrice
		} else {
			usdt_price = 0
			cny_price = 0
		}
		wallet_stat_db_lst, _ := models.GetWalletStatWid(1)
		var wallet_dstat_list []wallet.AddressStatData
		for _, wallet_stat := range wallet_stat_db_lst {
			w_sdata_ := wallet.AddressStatData{
				Amount: wallet_stat.Amount,
				DateTime:  wallet_stat.CreatedAt.Format("2006-01-02 15:04:05"),
				Time: wallet_stat.CreatedAt.Format("15:04"),
			}
			wallet_dstat_list = append(wallet_dstat_list, w_sdata_)
		}
		data := map[string]interface{}{
			"balance": balance_h,
			"icon": m_asst.Icon,
			"name": m_asst.Name,
			"chain_name": m_asst.ChainName,
			"adddress": address_balance.Address,
			"usdt_price": usdt_price,
			"cny_price": cny_price,
			"unit": 18,
			"data_stat": wallet_dstat_list,
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "wallet success")
		this.ServeJSON()
		return
	}
}


// GetWalletBalance @Title GetWalletBalance
// @Description 获取钱包余额 GetWalletBalance
// @Success 200 status bool, data interface{}, msg string
// @router /get_wallet_balance [post]
func (this *WalletController) GetWalletBalance() {
	var wallet_balance wallet.WalletBalanceReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &wallet_balance); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := wallet_balance.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	add_lst, _ := models.GetWalletBalance(wallet_balance.WalletUuid, wallet_balance.DeviceId)
	var wat_list []*wallet.WalletBalanceResponse
	var total_usd_asset float64
	if add_lst != nil {
		for _, value := range add_lst {
			asst := models.Asset{}
			m_asset := models.MarketAsset{}
			_ = orm.NewOrm().QueryTable(models.Asset{}).Filter("id", value.AssetId).One(&asst)
			_ = orm.NewOrm().QueryTable(models.MarketAsset{}).Filter("name", asst.Name).One(&m_asset)
			balance_h := math.Trunc(value.Balance*1e4+0.5) * 1e-4
			var usdt_price,  cny_price float64
			market_p := models.GetMarketByAssetId(m_asset.Id)
			if market_p != nil {
				usdt_price = balance_h * market_p.UsdPrice
				cny_price = balance_h * market_p.CnyPrice
			} else {
				usdt_price = 0
				cny_price = 0
			}
			total_usd_asset += usdt_price
			wat := &wallet.WalletBalanceResponse{
				Id: value.Id,
				DeviceId: value.DeviceId,
				WalletUuid: value.WalletUuid,
				Balance: balance_h,
				Icon: asst.Icon,
				Name: asst.Name,
				WalletName: value.WalletName,
				ContractAddr: value.ContractAddr,
				ChainName: "Ethereum",
				Address: value.Address,
				UsdtPrice: usdt_price,
				CnyPrice:  cny_price,
			}
			wat_list = append(wat_list, wat)
		}
	}
	data := map[string]interface{}{
		"total_asset": total_usd_asset,
		"coin_asset": wat_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "wallet success")
	this.ServeJSON()
	return
}


// GetTxByAddress @Title GetTxByAddress
// @Description 根据地址查询交易记录 GetTxByAddress
// @Success 200 status bool, data interface{}, msg string
// @router /get_tx_by_address [post]
func (this *WalletController) GetTxByAddress() {
	var address_record wallet.WalletRecordByAddrReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &address_record); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := address_record.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	ethersan_rcd := eth.BaseParamData {
		Url: EtherScanBaseUrl,
		Module: "account",
		Action: address_record.Action,
		Address: address_record.Address,
		Contractaddress:address_record.ContractAddress,
		Startblock: Startblock,
		Endblock: Endblock,
		Page: address_record.Page,
		Offset: address_record.PageSize,
		Sort: "desc",
		Apikey: ApiKey,
	}
	if address_record.Action == "txlist" {
		tx_data, err := ethersan_rcd.GetErc20TxByAddress()
		if err != nil {
			this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
			this.ServeJSON()
			return
		}
		var record_data_list []*wallet.TransRecordDateRep
		for _, value := range tx_data.Result {
			times_int, _ := strconv.ParseInt(value.TimeStamp, 10, 64)
			tm := time.Unix(times_int, 0)
			var in_out string
			if address_record.Address == value.From {
				in_out = "from"
			} else {
				in_out = "to"
			}
			record_data := &wallet.TransRecordDateRep{
				BlockNumber: value.BlockNumber,
				DateTine: tm.Format("2006-01-02 15:04:05"),
				AssetName: "ETH",
				Hash: value.Hash,
				From: value.From,
				To: value.To,
				Value: value.Value,
				ContractAddress: value.ContractAddress,
				GasUsed: value.GasUsed,
				GasPrice: value.GasPrice,
				IsError: value.IsError,
				TxreceiptStatus: value.TxreceiptStatus,
				TxInOut: in_out,
				Unit: 18,
			}
			record_data_list = append(record_data_list, record_data)
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, record_data_list, "wallet success")
		this.ServeJSON()
		return
	} else {
		tx_erc20_data, err := ethersan_rcd.GetTxByAddress()
		if err != nil {
			this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
			this.ServeJSON()
			return
		}
		var record_data_list []*wallet.TransRecordDateRep
		for _, value := range tx_erc20_data.Result {
			if value.ContractAddress == address_record.ContractAddress {
				times_int, _ := strconv.ParseInt(value.TimeStamp, 10, 64)
				tm := time.Unix(times_int, 0)
				var in_out string
				if address_record.Address == value.From {
					in_out = "from"
				} else {
					in_out = "to"
				}
				token_config := models.TokenConfig{}
				_ = orm.NewOrm().QueryTable(models.TokenConfig{}).Filter("contract_addr", value.ContractAddress).One(&token_config)
				record_data := &wallet.TransRecordDateRep {
					BlockNumber: value.BlockNumber,
					DateTine: tm.Format("2006-01-02 15:04:05"),
					AssetName: token_config.TokenSymbol,
					Hash: value.Hash,
					From: value.From,
					To: value.To,
					Value: value.Value,
					ContractAddress: value.ContractAddress,
					GasUsed: value.GasUsed,
					GasPrice: value.GasPrice,
					IsError: "0",
					TxreceiptStatus: "1",
					TxInOut: in_out,
					Unit: token_config.Decimal,
				}
				record_data_list = append(record_data_list, record_data)
			}
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, record_data_list, "wallet success")
		this.ServeJSON()
		return
	}
}


// GetSignTxInfo @Title GetSignTxInfo
// @Description 获取签名需要的信息，包含 Nonce, gas等 GetSignTxInfo
// @Success 200 status bool, data interface{}, msg string
// @router /get_sign_tx_info [post]
func (this *WalletController) GetSignTxInfo() {
	var sign_data wallet.SignDataRequest
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &sign_data); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := sign_data.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	nonce, _ := ethClient_.GetNonce(sign_data.Address)
	gs_list := models.GetGasNow()
	var gaslst []wallet.GasPriceResponse
	if gs_list != nil {
		for _, value := range gs_list {
			gs := wallet.GasPriceResponse{
				Index: value.Index,
				GasPrice: value.GasPrice,
			}
			gaslst = append(gaslst, gs)
		}
	}
	m_asset := models.MarketAsset{}
	market := models.Market{}
	_ = orm.NewOrm().QueryTable(models.MarketAsset{}).Filter("name", "ETH").One(&m_asset)
	_ = orm.NewOrm().QueryTable(models.Market{}).Filter("mk_asset_id", m_asset.Id).One(&market)
	data := map[string]interface{}{
		"usdt_pirce": market.UsdPrice,
		"nonce": nonce,
		"gas_limit": 90000,
		"gas_list": gs_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "wallet success")
	this.ServeJSON()
	return
}

// GetWalletAsset @Title GetWalletAsset
// @Description 钱包资产预览 GetWalletAsset
// @Success 200 status bool, data interface{}, msg string
// @router /get_wallet_asset [post]
func (this *WalletController) GetWalletAsset() {
	wat_asst := wallet.WalletAssetReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &wat_asst); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := wat_asst.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	var wat_asst_list []*wallet.WalletAssetRep
	var total_usd_asset float64
	name_wallet_dist, _ := models.GetWalletByDeviceId(wat_asst.DeviceId)
	var wallet_repeat_name_list []string
	for _, name_wallet := range name_wallet_dist{
		wallet_repeat_name_list =append(wallet_repeat_name_list, name_wallet.WalletName)
	}
	wallet_name_list := RemoveRepeatedElement(wallet_repeat_name_list)
	for _, w_name := range wallet_name_list {
		wallet_lst, _ := models.GetWalletBalanceName(w_name, wat_asst.DeviceId)
		var wbr_list []*wallet.WalletBalanceResponse
		if wallet_lst != nil {
			for _, value := range wallet_lst {
				asst := models.Asset{}
				m_asset := models.MarketAsset{}
				_ = orm.NewOrm().QueryTable(models.Asset{}).Filter("id", value.AssetId).One(&asst)
				_ = orm.NewOrm().QueryTable(models.MarketAsset{}).Filter("name", asst.Name).One(&m_asset)
				balance_h := math.Trunc(value.Balance*1e4+0.5) * 1e-4
				var usdt_price,  cny_price float64
				market_p := models.GetMarketByAssetId(m_asset.Id)
				if market_p != nil {
					usdt_price = balance_h * market_p.UsdPrice
					cny_price = balance_h * market_p.CnyPrice
				} else {
					usdt_price = 0
					cny_price = 0
				}
				total_usd_asset += usdt_price
				wat_balance := &wallet.WalletBalanceResponse{
					Id: value.Id,
					Balance: balance_h,
					Icon: asst.Icon,
					Name: asst.Name,
					ChainName: "Ethereum",
					UsdtPrice: usdt_price,
					CnyPrice:  cny_price,
				}
				wbr_list = append(wbr_list, wat_balance)
			}
		}
		wat := &wallet.WalletAssetRep{
			WalletName: w_name,
			WalletBalance: wbr_list,
		}
		wat_asst_list = append(wat_asst_list, wat)
	}
	data := map[string]interface{}{
		"total_asset": total_usd_asset,
		"coin_asset": wat_asst_list,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "wallet success")
	this.ServeJSON()
	return
}

func newIntegerFromString(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 0)
}

func mul(i *big.Int, i2 *big.Int) *big.Int {
	return new(big.Int).Mul(i, i2)
}

func (i Int) bigInt() *big.Int {
	return new(big.Int).Set(i.i)
}

func NewIntWithDecimal(n int64, dec int) Int {
	if dec < 0 {
		panic("NewIntWithDecimal() decimal is negative")
	}
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(dec)), nil)
	i := new(big.Int)
	i.Mul(big.NewInt(n), exp)
	if i.BitLen() > maxBitLen {
		panic("NewIntWithDecimal() out of bound")
	}
	return Int{i}
}


func FloatToBigInt(val float64, dec int64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	coin := new(big.Float)
	if dec == 0 {
		coin.SetInt(big.NewInt(1000000000000000000))
	} else {
		coin.SetInt(big.NewInt(dec))
	}
	bigval.Mul(bigval, coin)
	result := new(big.Int)
	bigval.Int(result)
	return result
}


// SendTx @Title SendTx
// @Description 发送交易到区块链网 SendTx
// @Success 200 status bool, data interface{}, msg string
// @router /send_tx [post]
func (this *WalletController) SendTx() {
	tx_bd := wallet.BroadCastTxReq{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &tx_bd); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := tx_bd.ParamCheck(); err != nil {
		fmt.Println(err)
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	address_entity := models.GetWalletSignInfo(tx_bd.DeviceId, tx_bd.WalletUuid, tx_bd.ContractAddr, tx_bd.AssetName)
	if address_entity != nil {
		var data []byte
		var to_address string
		var amt_lst *big.Int
		if tx_bd.ContractAddr != "" {
			amt_lst_erc20 := FloatToBigInt(tx_bd.Amount, 10e5)
			erc20_data, err := eth.ABITransfer(tx_bd.ToAddress, amt_lst_erc20)
			if err != nil {
				fmt.Println(err)
			}
			data = erc20_data
			to_address = tx_bd.ContractAddr
			amt_lst = big.NewInt(0)
		} else {
			data = []byte("")
			amt_lst = FloatToBigInt(tx_bd.Amount, 0)
			to_address = tx_bd.ToAddress
		}
		nonce, _ := ethClient_.GetNonce(tx_bd.FromAddress)
		big_int_gprice := new(big.Int)
		big_int_gprice, _ = big_int_gprice.SetString(tx_bd.GasPrice, 10)
		unsign_data := eth.EtherTx{
			ChainId:   1,
			PrvKeyStr: address_entity.PrivateKey,
			Nonce: int64(nonce),
			Address: to_address,
			Amount:  amt_lst,
			GasLimit: GasLimit,
			GasPrice: big_int_gprice,
			Data: hex.EncodeToString(data),
		}
		sign_tx_hex, err := eth.SignEthTx(unsign_data)
		if err != nil {
			fmt.Println(err)
			this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
			this.ServeJSON()
			return
		}
		tx_hash, err := ethClient_.SendTxHex(sign_tx_hex)
		if err != nil {
			this.Data["json"] = RetResource(false, types.HandleError, nil, err.Error())
			this.ServeJSON()
			return
		}
		data_ret := map[string]interface{}{
			"tx_hex": sign_tx_hex,
			"tx_hash": tx_hash,
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data_ret, "wallet success")
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	}
}

