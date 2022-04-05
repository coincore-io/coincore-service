package services

import (
	beego_pagination "coinwallet/common/utils/beego-pagination"
	"coinwallet/models"
	"github.com/astaxie/beego/orm"
	"net/url"
)

type RecordService struct {
	BaseService
}

func (Self *RecordService) GetPaginateDataWithDraw(listRows int, params url.Values) ([]*models.TxRecordList, beego_pagination.Pagination) {
	var data []*models.TxRecordList
	var total int64
	om := orm.NewOrm()
	inner := "from  tx_record as t0 inner join asset as t1 on t1.id = t0.asset_id  where t0.id > 0 "
	sql := "select t0.*,t1.name as asset_name " + inner
	sql1 := "select count(*) total " + inner

	//搜索、查询字段赋值
	Self.SearchField = append(Self.SearchField, new(models.TxRecord).SearchField()...)
	where,param := Self.ScopeWhereRaw(params)
	Self.PaginateRaw(listRows,params)

	if err := om.Raw(sql1+where,param).QueryRow(&total);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	Self.Pagination.Total = int(total)
	param = append(param,listRows*(Self.Pagination.CurrentPage-1),listRows)
	if _,err := om.Raw(sql+where+" order by created_at desc limit ?,?",param).QueryRows(&data);err != nil {
		return nil,beego_pagination.Pagination{}
	}
	return data,Self.Pagination
}