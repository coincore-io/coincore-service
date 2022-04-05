package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/orm"
	"time"
)

type Chain struct {
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"链ID" json:"id"`
	Name           string     `orm:"column(name);unique;index" description:"链名称" json:"name"`
	Mark           string     `orm:"column(mark);unique;index" description:"链标识" json:"mark"`
	Icon           string     `orm:"column(icon)" description:"图片ICON" json:"icon"`
	IsRemoved      int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt      time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt      time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *Chain) TableName() string {
	return common.TableName("chain")
}

func (this *Chain) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *Chain) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Chain) SearchField() []string {
	return []string{"name"}
}


func GetConfigList() ([]*Chain) {
	var chain_list []*Chain
	_, err := orm.NewOrm().QueryTable(&Chain{}).Filter("is_removed", 0).All(&chain_list)
	if err != nil {
		return nil
	}
	return chain_list
}

