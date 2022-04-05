package api_v1

import (
	"coinwallet/models"
	"coinwallet/types"
	"github.com/astaxie/beego"
)

type ConfigController struct {
	beego.Controller
}

// GetConfigList @Title GetConfigList
// @Description 获取默认配置 GetConfigList
// @Success 200 status bool, data interface{}, msg string
// @router /get_config_list [post]
func (this *ConfigController) GetConfigList() {
	chain_list  := models.GetConfigList()
	if chain_list == nil {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, chain_list, "wallet success")
	this.ServeJSON()
	return
}
