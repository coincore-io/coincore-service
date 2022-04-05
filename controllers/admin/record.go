package controllers

import "coinwallet/services"

type RecordController struct {
	baseController
}

func (Self *RecordController) Index() {
	var srv services.RecordService
	data, pagination := srv.GetPaginateDataWithDraw(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "record/index.html"
}