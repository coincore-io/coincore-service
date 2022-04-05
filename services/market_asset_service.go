package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/form_validate"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type MarketAssetService struct {
	BaseService
}

func (self *MarketAssetService) GetPaginateData(listRows int, params url.Values) ([]*models.MarketAsset, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.MarketAsset).SearchField()...)

	var data []*models.MarketAsset
	o := orm.NewOrm().QueryTable(new(models.MarketAsset))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}

func (*MarketAssetService) Create(form *form_validate.MarketAssetForm) int {
	data := models.MarketAsset{
		Name: form.Name,
		Icon: form.Icon,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*MarketAssetService) GetById(id int64) *models.MarketAsset {
	o := orm.NewOrm()
	data := models.MarketAsset{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*MarketAssetService) Update(form *form_validate.MarketAssetForm) int{
	o := orm.NewOrm()
	data := models.MarketAsset{Id: form.Id}
	if o.Read(&data) == nil {
		data.Name = form.Name
		if len(form.Icon) > 0 {
			data.Icon = form.Icon
		}
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*MarketAssetService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.MarketAsset)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}

func (*MarketAssetService) ActiveAsset() []*models.MarketAsset{
	var data []*models.MarketAsset
	_, err := orm.NewOrm().QueryTable(new(models.MarketAsset)).Filter("is_removed__eq", 0).All(&data)
	if err == nil {
		return nil
	} else {
		return data
	}
}