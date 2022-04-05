package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/form_validate"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type NotebookAddr struct {
	BaseService
}

func (self *NotebookAddr) GetPaginateData(listRows int, params url.Values) ([]*models.AddrNoteBook, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.AddrNoteBook).SearchField()...)

	var data []*models.AddrNoteBook
	o := orm.NewOrm().QueryTable(new(models.AddrNoteBook))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}

func (*NotebookAddr) Create(form *form_validate.NotebookForm) int {
	data := models.AddrNoteBook{
		DeviceId: form.DeviceId,
		Name: form.Name,
		AssetName: form.AssetName,
		Memo: form.Memo,
		Addr: form.Addr,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*NotebookAddr) GetById(id int64) *models.AddrNoteBook {
	o := orm.NewOrm()
	data := models.AddrNoteBook{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*NotebookAddr) Update(form *form_validate.NotebookForm) int{
	o := orm.NewOrm()
	data := models.AddrNoteBook{Id: form.Id}
	if o.Read(&data) == nil {
		data.DeviceId = form.DeviceId
		data.Name = form.Name
		data.AssetName = form.AssetName
		data.Memo = form.Memo
		data.Addr = form.Addr
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*NotebookAddr) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.AddrNoteBook)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}