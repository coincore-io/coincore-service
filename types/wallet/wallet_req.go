package wallet

import (
	"coinwallet/types"
	"github.com/pkg/errors"
)

type SubmitWalletInfoReq struct {
	DeviceId     string `json:"device_id"`
	WalletUuid   string `json:"wallet_uuid"`
	WalletName   string `json:"wallet_name"`
	AssetName    string `json:"asset_name"`
	ChainName    string `json:"chain_name"`
	Address      string `json:"address"`
	ContractAddr string `json:"contract_addr"`
	WordCode     string `json:"word_code"`
	PrivateKey   string `json:"private_key"`
}

func (this SubmitWalletInfoReq) ParamCheck() (int, error) {
	if this.WalletUuid == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.AssetName == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.ChainName == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.Address == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.WordCode == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.PrivateKey == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}

type BatchSubmitWalletReq struct {
	BatchWallet  []SubmitWalletInfoReq `json:"batch_wallet"`
}

func (this BatchSubmitWalletReq) ParamCheck() (int, error) {
	if this.BatchWallet == nil {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}


type DeleteWalletReq struct {
	DeviceId     string `json:"device_id"`
	WalletUuid   string `json:"wallet_uuid"`
}

func (this DeleteWalletReq) ParamCheck() (int, error) {
	if this.DeviceId == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.WalletUuid == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}


type AddressBalanceReq struct {
	DeviceId     string `json:"device_id"`
	WalletUuid   string `json:"wallet_uuid"`
	WalletName   string `json:"wallet_name"`
	AssetName    string `json:"asset_name"`
	ChainName    string `json:"chain_name"`
	Address      string `json:"address"`
	ContractAddr string `json:"contract_addr"`
}

func (this AddressBalanceReq) ParamCheck() (int, error) {
	if this.WalletUuid == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.AssetName == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.ChainName == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.Address == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}


type WalletBalanceReq struct {
	DeviceId     string `json:"device_id"`
	WalletUuid   string `json:"wallet_uuid"`
	ChainName    string `json:"chain_name"`
}

func (this WalletBalanceReq) ParamCheck() (int, error) {
	if this.WalletUuid == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.DeviceId == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}


type WalletRecordByAddrReq struct {
	Action          string  `json:"action"`
	Address         string  `json:"address"`
	ContractAddress string  `json:"contract_address"` // 是ERC20传合约地址，不是传空
	Page            string  `json:"page"`
	PageSize        string  `json:"page_size"`
}

func (this WalletRecordByAddrReq) ParamCheck() (int, error) {
	if this.Address == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.Page == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.PageSize == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}


type SignDataRequest struct {
	ChainName   string `json:"chain_name"`
	Address     string `json:"address"`
}

func (this SignDataRequest) ParamCheck() (int, error) {
	if this.ChainName == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}

type WalletAssetReq struct {
	DeviceId     string   `json:"device_id"`
}


func (this WalletAssetReq) ParamCheck() (int, error) {
	if this.DeviceId == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}


type BroadCastTxReq struct {
	DeviceId     string  `json:"device_id"`
	WalletUuid   string  `json:"wallet_uuid"`
	FromAddress  string  `json:"from_address"`  // 转出账户
	ToAddress    string  `json:"to_address"`    // 转入账户
	Amount       float64 `json:"amount"`        // 转账金额
	AssetName    string  `json:"asset_name"`    // 转账币种
	ContractAddr string  `json:"contract_addr"` // 合约地址
	GasPrice     string  `json:"gas_price"`
	GasLimit     string  `json:"gas_limit"`
}

func (this BroadCastTxReq) ParamCheck() (int, error) {
	if this.FromAddress == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.ToAddress == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}


type DeleteWalletTokenReq struct {
	DeviceId     string `json:"device_id"`
	WalletUuid   string `json:"wallet_uuid"`
	ContractAddr string `json:"contract_addr"`
}

func (this DeleteWalletTokenReq) ParamCheck() (int, error) {
	if this.DeviceId == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.WalletUuid == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	if this.ContractAddr == "" {
		return types.ParamEmptyError, errors.New("wallet error")
	}
	return types.ReturnSuccess, nil
}