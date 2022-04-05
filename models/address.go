package models

import (
	"coinwallet/common"
	"coinwallet/types/wallet"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Address struct {
	Id           int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	AssetId      int64      `orm:"column(asset_id);size(11)" description:"资产ID" json:"asset_id"`
	DeviceId     string     `orm:"column(device_id);size(512)" description:"设备ID" json:"device_id"`
	WalletUuid   string     `orm:"column(wallet_uuid);size(512)" description:"钱包的UUID" json:"wallet_uuid"`
	WalletName   string     `orm:"column(wallet_name);size(512)" description:"钱包名称" json:"wallet_name"`
	Address      string     `orm:"column(address);size(512)" description:"地址" json:"address"`
	ContractAddr string     `orm:"column(contract_addr);default('')" description:"合约地址" json:"contract_addr"`
	PrivateKey   string     `orm:"column(private_key);size(512)" description:"加密的私钥" json:"private_key"`
	Balance      float64    `orm:"column(balance);default(1);digits(22);decimals(8)" description:"地址余额度" json:"balance"`
	IsRemoved    int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt    time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt    time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *Address) TableName() string {
	return common.TableName("address")
}

func (this *Address) SearchField() []string {
	return []string{"wallet_name"}
}

func (this *Address) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *Address) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *Address) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *Address) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func SubmitWalletInfo(req wallet.SubmitWalletInfoReq) error {
	var err error
	var chain Chain
	err = orm.NewOrm().QueryTable(Chain{}).
		Filter("name", req.ChainName).
		OrderBy("-id").OrderBy("-id").One(&chain)
	if err != nil {
		return err
	}
	var asset Asset
	err = orm.NewOrm().QueryTable(Asset{}).
		Filter("name", req.AssetName).
		Filter("chaind_id", chain.Id).
		OrderBy("-id").OrderBy("-id").One(&asset)
	if err == nil {
		ok := orm.NewOrm().QueryTable(Address{}).Filter("device_id",  req.DeviceId).
			Filter("wallet_uuid", req.WalletUuid).
			Filter("asset_id", asset.Id).
			Filter("contract_addr", req.ContractAddr).Exist()
		if ok == false {
			address := Address{
				DeviceId: req.DeviceId,
				AssetId: asset.Id,
				WalletUuid: req.WalletUuid,
				WalletName: req.WalletName,
				Address: req.Address,
				PrivateKey: req.PrivateKey,
				ContractAddr: req.ContractAddr,
			}
			err, _ := address.Insert()
			if err != nil {
				return err
			}
		} else {
			var address Address
			err = orm.NewOrm().QueryTable(Address{}).Filter("device_id",  req.DeviceId).
				Filter("wallet_uuid", req.WalletUuid).
				Filter("asset_id", asset.Id).
				Filter("contract_addr", req.ContractAddr).One(&address)
			if err != nil {
				return err
			}
			address.IsRemoved = 0
			err = address.Update()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func UpdateBalanceByAddress(asset_name, chain_name, deviced_id, wallet_uuid, contract_addr string, balance float64) {
	var err error
	var chain Chain
	err = orm.NewOrm().QueryTable(Chain{}).
		Filter("name", chain_name).
		OrderBy("-id").OrderBy("-id").One(&chain)
	if err != nil {
		fmt.Println(err)
	}
	var asset Asset
	err = orm.NewOrm().QueryTable(Asset{}).
		Filter("name", asset_name).
		Filter("chaind_id", chain.Id).
		OrderBy("-id").OrderBy("-id").One(&asset)
	if err != nil {
		fmt.Println(err)
	}
	addr := Address{}
	err = orm.NewOrm().QueryTable(Address{}).Filter("device_id", deviced_id).
		Filter("wallet_uuid", wallet_uuid).
		Filter("asset_id", asset.Id).
		Filter("contract_addr", contract_addr).
		One(&addr)
	if err != nil {
		fmt.Println(err)
	}
	addr.Balance = balance
	err = addr.Update()
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteWallet(deviced_id, wallet_uuid string) bool {
	var err error
	var addr_list []Address
	_, err = orm.NewOrm().QueryTable(Address{}).
		Filter("device_id", deviced_id).
		Filter("wallet_uuid", wallet_uuid).All(&addr_list)
	for _, addr := range addr_list {
		addr.IsRemoved = 1
		err = addr.Update()
		if err != nil {
			return false
		}
	}
	return true
}

func DeleteWalletToken(deviced_id, wallet_uuid, contract_addr string) bool {
	var err error
	var addr_list []Address
	_, err = orm.NewOrm().QueryTable(Address{}).
		Filter("device_id", deviced_id).
		Filter("wallet_uuid", wallet_uuid).
		Filter("contract_addr", contract_addr).All(&addr_list)
	for _, addr := range addr_list {
		addr.IsRemoved = 1
		err = addr.Update()
		if err != nil {
			return false
		}
	}
	return true
}

func GetWalletBalance(wallet_uuid string, device_id string) ([]*Address, error) {
	var balance_list []*Address
	_, err := orm.NewOrm().QueryTable(&Address{}).Filter("is_removed", 0).Filter("device_id", device_id).
		Filter("wallet_uuid", wallet_uuid).All(&balance_list)
	if err != nil {
		return nil, err
	}
	return balance_list, nil
}


func GetWalletByDeviceId(device_id string) ([]*Address, error) {
	var balance_list []*Address
	_, err := orm.NewOrm().QueryTable(&Address{}).Filter("is_removed", 0).Filter("device_id", device_id).All(&balance_list)
	if err != nil {
		return nil, err
	}
	return balance_list, nil
}


func GetWalletBalanceName(wallet_name string, device_id string) ([]*Address, error) {
	var balance_list []*Address
	_, err := orm.NewOrm().QueryTable(&Address{}).Filter("is_removed", 0).Filter("device_id", device_id).
		Filter("wallet_name", wallet_name).All(&balance_list)
	if err != nil {
		return nil, err
	}
	return balance_list, nil
}

func GetWalletSignInfo(device_id string, wallet_uuid string, contract_addr string, asset_name string) *Address  {
	var asset Asset
	err := orm.NewOrm().QueryTable(&Asset{}).Filter("Name", asset_name).OrderBy("-id").One(&asset)
	if err != nil {
		return nil
	}
	var address Address
	err = orm.NewOrm().QueryTable(&Address{}).
		Filter("is_removed", 0).
		Filter("device_id", device_id).
		Filter("wallet_uuid", wallet_uuid).
		Filter("contract_addr", contract_addr).Filter("asset_id", asset.Id).OrderBy("-id").One(&address)
	if err != nil {
		return nil
	}
	return &address
}


