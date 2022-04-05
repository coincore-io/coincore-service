package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/form_validate"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type ChainService struct {
	BaseService
}

func (self *ChainService) GetPaginateData(listRows int, params url.Values) ([]*models.Chain, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.Chain).SearchField()...)

	var data []*models.Chain
	o := orm.NewOrm().QueryTable(new(models.Chain))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}

func (*ChainService) Create(form *form_validate.ChainForm) int {
	data := models.Chain{
		Name: form.Name,
		Mark: form.Mark,
		Icon: form.Icon,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*ChainService) GetById(id int64) *models.Chain {
	o := orm.NewOrm()
	data := models.Chain{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*ChainService) Update(form *form_validate.ChainForm) int{
	o := orm.NewOrm()
	data := models.Chain{Id: form.Id}
	if o.Read(&data) == nil {
		data.Name = form.Name
		data.Mark = form.Mark
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

func (*ChainService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.Chain)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}

func (*ChainService) Chains() []*models.Chain{
	var data []*models.Chain
	_, err := orm.NewOrm().QueryTable(new(models.Chain)).Filter("is_removed", 0).All(&data)
	if err != nil {
		return nil
	} else {
		return data
	}
}