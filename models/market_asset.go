package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type MarketAsset struct {
	Id        int64      `orm:"pk;column(id);auto;size(11)" description:"币种ID" json:"id"`
	ChainName string     `orm:"column(chain_name)" description:"链名称" json:"chain_name"`
	Name      string     `orm:"column(name)" description:"币种名称" json:"name"`
	Icon      string     `orm:"column(icon)" description:"图片ICON" json:"icon"`
	IsRemoved int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *MarketAsset) TableName() string {
	return common.TableName("market_asset")
}

func(this *MarketAsset) SearchField() []string {
	return []string{"name"}
}

func (this *MarketAsset) Insert() error {
	if _, err := orm.NewOrm().Insert(this); err != nil {
		return err
	}
	return nil
}

func (this *MarketAsset) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func GetMarketAssetById(asset_id int64) (*MarketAsset, string, error) {
	var mkt_ast MarketAsset
	err := orm.NewOrm().QueryTable(MarketAsset{}).Filter("id", asset_id).One(&mkt_ast)
	if err != nil {
		return nil, "获取资产失败", errors.New( "获取资产失败")
	}
	return &mkt_ast, "", nil
}