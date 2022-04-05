package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/form_validate"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type AddressService struct {
	BaseService
}


func (self *AddressService) GetPaginateData(listRows int, params url.Values) ([]*models.Address, beego_pagination.Pagination) {
	//搜索、查询字段赋值
	self.SearchField = append(self.SearchField, new(models.Address).SearchField()...)

	var data []*models.Address
	o := orm.NewOrm().QueryTable(new(models.Address))
	_, err := self.PaginateAndScopeWhere(o, listRows, params).All(&data)
	if err != nil {
		return nil, self.Pagination
	} else {
		return data, self.Pagination
	}
}

func (*AddressService) Create(form *form_validate.AddressForm) int {
	//TODO
	return 0
}

func (*AddressService) GetById(id int64) *models.Address {
	o := orm.NewOrm()
	data := models.Address{Id: id}
	err := o.Read(&data)
	if err != nil {
		return nil
	}
	return &data
}


func (*AddressService) Update(form *form_validate.AddressForm) int{
	o := orm.NewOrm()
	data := models.Address{Id: form.Id}
	if o.Read(&data) == nil {
		data.AssetId = form.AssetId
		data.AssetId = form.AssetId
		data.DeviceId = form.DeviceId
		data.WalletName = form.WalletName
		data.WalletUuid = form.WalletUuid
		data.Address = form.Address
		data.ContractAddr = form.ContractAddr
		data.PrivateKey = form.PrivateKey
		data.Balance = form.Balance
		num, err := o.Update(&data)
		if err == nil {
			return int(num)
		} else {
			return 0
		}
	}
	return 0
}

func (*AddressService) Del(ids []int) int{
	count, err := orm.NewOrm().QueryTable(new(models.Address)).Filter("id__in", ids).Delete()
	if err == nil {
		return int(count)
	} else {
		return 0
	}
}
