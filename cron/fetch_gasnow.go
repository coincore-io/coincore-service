package cron

import (
	"coinwallet/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)


type gasDataItem struct {
	Index    int64 `json:"index"`
	GasPrice int64 `json:"gasPrice"`
}

type gasDatalist struct {
	List []*gasDataItem `json:"list"`
}

type gasNowData struct {
	Code int `json:"code"`
	Data gasDatalist `json:"data"`
}


const BaseUrl = "https://gasnow.sparkpool.com/api/v2/gas/price"


func SyncGasPrice() error {
	resp, _ := http.Get(BaseUrl)
	defer resp.Body.Close()
	ret_data, _ := ioutil.ReadAll(resp.Body)
	var gas_now_data gasNowData
	if err := json.Unmarshal([]byte(ret_data), &gas_now_data); err != nil {
		return errors.New("unmarshall json fail")
	}
	if gas_now_data.Code != 200 {
		fmt.Println("fetch gasnow data fail, code is", gas_now_data.Code)
		return errors.New("fetch gasnow data fail")
	}
	for _, value := range gas_now_data.Data.List {
		var gas_now_db models.GasNow
		err := orm.NewOrm().QueryTable(models.GasNow{}).Filter("index", value.Index).One(&gas_now_db)
		if err != nil {
			gas_now := models.GasNow{
				Index: value.Index,
				GasPrice: value.GasPrice,
			}
			err, _ := gas_now.Insert()
			if err != nil {
				return err
			}
		} else {
			gas_now_db.GasPrice = value.GasPrice
			err := gas_now_db.Update()
			if err != nil {
				return err
			}
		}

	}


	return nil
}
