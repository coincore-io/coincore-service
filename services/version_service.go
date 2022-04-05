package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/form_validate"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type VersionService struct {
	BaseService
}

func (self *VersionService) GetPaginateData(listRows int, params url.Values) ([]*models.Version, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.Version).SearchField()...)

	var data []*models.Version
	o := orm.NewOrm().QueryTable(new(models.Version))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}

func (*VersionService) Create(form *form_validate.VersionForm) int {
	data := models.Version{
		VersionNum: form.VersionNum,
		Platforms: form.Platforms,
		Decribe: form.Decribe,
		DownloadUrl: form.DownloadUrl,
		IsForce: form.IsForce,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*VersionService) GetById(id int64) *models.Version {
	o := orm.NewOrm()
	data := models.Version{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*VersionService) Update(form *form_validate.VersionForm) int{
	o := orm.NewOrm()
	data := models.Version{Id: form.Id}
	if o.Read(&data) == nil {
		data.VersionNum = form.VersionNum
		data.Platforms = form.Platforms
		data.Decribe = form.Decribe
		data.DownloadUrl = form.DownloadUrl
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*VersionService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.Version)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}