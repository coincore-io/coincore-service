package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type GasService struct {
	BaseService
}

func (self *GasService) GetPaginateData(listRows int, params url.Values) ([]*models.GasNow, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.GasNow).SearchField()...)

	var data []*models.GasNow
	o := orm.NewOrm().QueryTable(new(models.GasNow))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}