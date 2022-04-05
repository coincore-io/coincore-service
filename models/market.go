package models

import (
	"coinwallet/common"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Market struct {
	Id        int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	MkAssetId int64      `orm:"column(mk_asset_id)" description:"行情的资产ID" json:"mk_asset_id"`
	MarketAsset *MarketAsset   `orm:"rel(fk)"  json:"market_asset"`
	UsdPrice  float64    `orm:"column(usd_price)" description:"USD价格" json:"usd_price"`
	CnyPrice  float64    `orm:"column(cny_price)" description:"USD价格" json:"cny_price"`
	Rate      float64    `orm:"column(rate)" description:"涨幅" json:"rate"`
	IsRemoved int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *Market) TableName() string {
	return common.TableName("market")
}

func (this *Market) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *Market) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *Market) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *Market) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func GetMarketList(page, page_size int64) ([]*Market, int, string) {
	var nl []*Market
	filter := orm.NewOrm().QueryTable(&Market{})
	total, err := filter.Count()
	if err != nil {
		return nil, 0, "获取行情记录数量失败"
	}
	_, err = filter.Limit(page_size, page_size*(page-1)).OrderBy("-usd_price").All(&nl)
	if err != nil {
		return nil, 0, "获取行情列表失败"
	}
	return nl, int(total), "获取行情列表成功"
}

func GetMarketByAssetId(asset_id int64) *Market {
	var mkt_ast Market
	err := orm.NewOrm().QueryTable(Market{}).Filter("mk_asset_id", asset_id).One(&mkt_ast)
	if err != nil {
		return nil
	}
	return &mkt_ast
}

func GetMarketAssetIds() []*Market {
	var data []*Market
	orm.NewOrm().QueryTable("market").RelatedSel().All(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
