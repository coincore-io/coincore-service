package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/orm"
	"time"
)


type GasNow struct {
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	Index          int64      `orm:"column(index);default(0)" description:"Index" json:"index"`
	GasPrice       int64      `orm:"column(gas_price);default(0)" description:"gas_price" json:"gas_price"`
	IsRemoved      int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt      time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt      time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *GasNow) TableName() string {
	return common.TableName("gasnow")
}

func(this *GasNow) SearchField() []string {
	return []string{}
}

func (this *GasNow) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *GasNow) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}


func (this *GasNow) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *GasNow) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func GetGasNow() ([]*GasNow){
	var gas_list []*GasNow
	_, err := orm.NewOrm().QueryTable(&GasNow{}).All(&gas_list)
	if err != nil {
		return nil
	}
	return gas_list
}