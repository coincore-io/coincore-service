package api_v1

import (
	"coinwallet/models"
	"coinwallet/types"
	"coinwallet/types/market"
	"encoding/json"
	"github.com/astaxie/beego"
)


type MarketController struct {
	beego.Controller
}

// GetMarketPrice @Title GetMarketPrice
// @Description 获取币种行情 GetMarketPrice
// @Success 200 status bool, data interface{}, msg string
// @router /get_market_price [post]
func (this *MarketController) GetMarketPrice() {
	 pageP := types.PageSizeData{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &pageP); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	mkt_lst, total, msg := models.GetMarketList(int64(pageP.Page), int64(pageP.PageSize))
	var market_lists [] market.MarketPriceRep
	for _, value := range mkt_lst {
		ma, msg, err := models.GetMarketAssetById(value.MkAssetId)
		if err != nil {
			this.Data["json"] = RetResource(false, types.SystemDbErr, nil, msg)
			this.ServeJSON()
			return
		}
		news_r := market.MarketPriceRep{
			Id: value.Id,
			Name: ma.Name,
			ChainName: ma.ChainName,
			Icon: ma.Icon,
			UsdPrice: value.UsdPrice,
			CnyPrice: value.CnyPrice,
			Rate: value.Rate,
		}
		market_lists = append(market_lists, news_r)
	}
	data := map[string]interface{}{
		"total":     total,
		"gds_lst":   market_lists,
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data, msg)
	this.ServeJSON()
	return
}
