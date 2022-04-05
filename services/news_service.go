package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/form_validate"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type NewsService struct {
	BaseService
}

func (self *NewsService) GetPaginateData(listRows int, params url.Values) ([]*models.News, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.News).SearchField()...)

	var data []*models.News
	o := orm.NewOrm().QueryTable(new(models.News))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}

func (*NewsService) Create(form *form_validate.NewsForm) int {
	data := models.News{
		CatName: form.CatName,
		Content: form.Content,
		Author: form.Author,
		Abstract: form.Abstract,
		Title: form.Title,
		Image: form.Image,
		NewsType: form.NewsType,
	}
	id, err := orm.NewOrm().Insert(&data)
	if err == nil {
		return int(id)
	} else {
		return 0
	}
}

func (*NewsService) GetById(id int64) *models.News {
	o := orm.NewOrm()
	data := models.News{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*NewsService) Update(form *form_validate.NewsForm) int{
	o := orm.NewOrm()
	data := models.News{Id: form.Id}
	if o.Read(&data) == nil {
		data.CatName = form.CatName
		data.NewsType = form.NewsType
		data.Abstract = form.Abstract
		data.Author = form.Author
		if len(form.Image) > 0 {
			data.Image = form.Image
		}
		data.Content = form.Content
		data.Title = form.Title
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*NewsService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.News)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}