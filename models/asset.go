package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/orm"
	"time"
)


type Asset struct {
	Id             int64      `orm:"pk;column(id);auto;size(11)" description:"币种ID" json:"id"`
	ChaindId       int64      `orm:"column(chaind_id);size(11)" description:"链ID" json:"chaind_id"`
	Name           string     `orm:"column(name);size(11);index" description:"币种名称" json:"name"`
	Icon           string     `orm:"column(icon);unique;index" description:"图片ICON" json:"icon"`
	Unit           int64      `orm:"column(unit);default(8)" description:"币种精度" json:"unit"`
	IsRemoved      int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt      time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt      time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *Asset) TableName() string {
	return common.TableName("asset")
}

func (this *Asset) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *Asset) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (this *Asset) SearchField() []string {
	return []string{"name"}
}

