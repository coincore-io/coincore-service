package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type MarketService struct {
	BaseService
}

func (self *MarketService) GetPaginateData(listRows int, params url.Values) ([]*models.Market, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.MarketAsset).SearchField()...)

	var data []*models.Market
	o := orm.NewOrm().QueryTable(new(models.Market)).RelatedSel()
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}


