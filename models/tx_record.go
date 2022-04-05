package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/orm"
	"time"
)

type TxRecord struct {
	Id           int64     `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	AssetId      int64     `orm:"column(asset_id);size(11)" description:"资产ID" json:"asset_id"`
	FromAddr     string    `orm:"column(from_addr);size(512)" description:"发送方" json:"from_addr"`
	ToAddr    	 string    `orm:"column(to_addr);size(512)" description:"接收方" json:"to_addr"`
	Amount       string    `orm:"column(amount);size(512)" description:"转账金额" json:"amount"`
	Memo    	 string    `orm:"column(memo);size(512)" description:"备注" json:"memo"`
	Hash         string    `orm:"column(hash);size(512)" description:"交易Hash" json:"hash"`
	BlockHeight  int64     `orm:"column(block_height)" description:"所在区块" json:"block_height"`
	TxTime       string    `orm:"column(tx_time);size(512)" description:"交易时间" json:"tx_time"`
	IsRemoved    int8      `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt    time.Time `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt    time.Time `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

type TxRecordList struct {
	TxRecord
	AssetName				string				`json:"asset_name"`
}

func (this *TxRecord) SearchField() []string {
	return []string{}
}

func (this *TxRecord) TableName() string {
	return common.TableName("tx_record")
}

func (this *TxRecord) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *TxRecord) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *TxRecord) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *TxRecord) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

