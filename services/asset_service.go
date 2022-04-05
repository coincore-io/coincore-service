package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/form_validate"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type AssetService struct {
	BaseService
}

func (self *AssetService) GetPaginateData(listRows int, params url.Values) ([]*models.Asset, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.Asset).SearchField()...)

	var data []*models.Asset
	o := orm.NewOrm().QueryTable(new(models.Asset))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}

func (*AssetService) Create(form *form_validate.AssetForm) int {
	data := models.Asset{
		Name: form.Name,
		ChaindId: form.Uni,
		Icon: form.Icon,
		Unit: form.Unit,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*AssetService) GetById(id int64) *models.Asset {
	o := orm.NewOrm()
	data := models.Asset{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*AssetService) Update(form *form_validate.AssetForm) int{
	o := orm.NewOrm()
	data := models.Asset{Id: form.Id}
	if o.Read(&data) == nil {
		data.Name = form.Name
		data.ChaindId = form.Uni
		data.Unit = form.Unit
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

func (*AssetService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.Asset)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}

func (*AssetService) ActiveAsset() []*models.Asset{
	var data []*models.Asset
	_, err := orm.NewOrm().QueryTable(new(models.Asset)).Filter("is_removed", 0).All(&data)
	if err != nil {
		return nil
	} else {
		return data
	}
}