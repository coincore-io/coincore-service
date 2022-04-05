package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/orm"
	"time"
)

type WalletStat struct {
	Id           int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	WalletId     int64      `orm:"column(wallet_id);size(11)" description:"钱包ID" json:"wallet_id"`
	Amount       float64    `orm:"column(amount);default(1);digits(22);decimals(8)" description:"额度变动" json:"amount"`
	IsRemoved    int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt    time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt    time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *WalletStat) TableName() string {
	return common.TableName("wallet_stat")
}

func (this *WalletStat) SearchField() []string {
	return []string{"wallet_name"}
}

func (this *WalletStat) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *WalletStat) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *WalletStat) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *WalletStat) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func GetWalletStatWid(wallet_id int64) ([]*WalletStat, error) {
	var wallet_stat_list []*WalletStat
	_, err := orm.NewOrm().QueryTable(&WalletStat{}).
		Filter("is_removed", 0).
		Filter("wallet_id", wallet_id).All(&wallet_stat_list)
	if err != nil {
		return nil, err
	}
	return wallet_stat_list, nil
}
