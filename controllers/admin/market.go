package controllers

import "coinwallet/services"

type MarketController struct {
	baseController
}

func (Self *MarketController) Index() {
	var srv services.MarketService
	data, pagination := srv.GetPaginateData(admin["per_page"].(int), gQueryParams)
	Self.Data["data"] = data
	Self.Data["paginate"] = pagination

	Self.Layout = "public/base.html"
	Self.TplName = "market/index.html"
}