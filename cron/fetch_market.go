package cron

import (
	"coinwallet/models"
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)


type MarketDataItem struct {
	CurrentPrice    float64  `json:"current_price"`
	CurrentPriceUsd float64 `json:"current_price_usd"`
	Code  string   `json:"code"`
	Name  string   `json:"name"`
	ChangePercent float64 `json:"change_percent"`
}

type MarketData struct {
	Maxpage int64 `json:"maxpage"`
	Currpage int64 `json:"currpage"`
	Code int64 `json:"code"`
	Msg string `json:"msg"`
	Data []MarketDataItem `json:"data"`
}

func RealMarketAssetPrice() (err error) {
	logs.Info("start exec RealMarketAssetPrice")
	db := orm.NewOrm()
	err = db.Begin()
	defer func() {
		if err != nil {
			err = db.Rollback()
			err = errors.Wrap(err, "rollback db transaction error in RealAssetPrice")
		} else {
			err = errors.Wrap(db.Commit(), "commit db transaction error in RealAssetPrice")
		}
	}()
	req_url := "https://dncapi.bqiapp.com/api/coin/web-coinrank?page=1&type=-1&pagesize=%s&webp=1"
	resp, _ := http.Get(req_url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var ret_data MarketData
	if err := json.Unmarshal(body, &ret_data); err != nil {
		return err
	}
	for _, value := range ret_data.Data {
		var market_asset models.MarketAsset
		_ = db.QueryTable(models.MarketAsset{}).Filter("name", value.Name).One(&market_asset)
		if market_asset.Name == value.Name {
			var market_price models.Market
			err = db.QueryTable(models.Market{}).Filter("mk_asset_id", market_asset.Id).One(&market_price)
			if err == nil {
				market_price.UsdPrice = value.CurrentPriceUsd
				market_price.CnyPrice = value.CurrentPrice
				market_price.Rate = value.ChangePercent
				_ = market_price.Update()
			} else {
				mp := models.Market{
					MkAssetId: market_asset.Id,
					MarketAsset: &market_asset,
					UsdPrice: value.CurrentPriceUsd,
					CnyPrice: value.CurrentPrice,
					Rate: value.ChangePercent,
				}
				_, _ = mp.Insert()
			}
		}
	}
	logs.Info("end exec RealMarketAssetPrice")
	return nil
}