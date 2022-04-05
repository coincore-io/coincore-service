package controllers

import "coinwallet/services"

type GasController struct {
	baseController
}

func (Self *GasController) Index() {
	var srv services.GasService
	data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "gas/index.html"
}