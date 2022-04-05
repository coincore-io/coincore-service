package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/form_validate"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type TokenConfigService struct {
	BaseService
}

func (self *TokenConfigService) GetPaginateData(listRows int, params url.Values) ([]*models.TokenConfig, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.TokenConfig).SearchField()...)

	var data []*models.TokenConfig
	o := orm.NewOrm().QueryTable(new(models.TokenConfig))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}

func (*TokenConfigService) Create(form *form_validate.TokenConfigForm) int {
	data := models.TokenConfig{
		AssetId: form.AssetId,
		TokenName: form.TokenName,
		Icon: form.Icon,
		TokenSymbol: form.TokenSymbol,
		ContractAddr: form.ContractAddr,
		Decimal: form.Decimal,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*TokenConfigService) GetById(id int64) *models.TokenConfig {
	o := orm.NewOrm()
	data := models.TokenConfig{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*TokenConfigService) Update(form *form_validate.TokenConfigForm) int{
	o := orm.NewOrm()
	data := models.TokenConfig{Id: form.Id}
	if o.Read(&data) == nil {
		data.AssetId = form.AssetId
		data.TokenName = form.TokenName
		data.TokenSymbol = form.TokenSymbol
		data.ContractAddr = form.ContractAddr
		if len(form.Icon) > 0 {
			data.Icon = form.Icon
		}
		data.Decimal = form.Decimal
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*TokenConfigService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.TokenConfig)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}