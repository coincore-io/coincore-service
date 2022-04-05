package api_v1

import (
	"coinwallet/models"
	"coinwallet/types"
	"encoding/json"
	"github.com/astaxie/beego"
)

type VersionController struct {
	beego.Controller
}

type VersionDate struct {
	Platforms    int64     `json:"platforms"`  // 0: 安卓 1: IOS
}

// GetVersionInfo @Title GetVersionInfo
// @Description 获取版本信息 GetVersionInfo
// @Success 200 status bool, data interface{}, msg string
// @router /version_info [post]
func (vc *VersionController) GetVersionInfo() {
	var vd VersionDate
	if err := json.Unmarshal(vc.Ctx.Input.RequestBody, &vd); err != nil {
		vc.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		vc.ServeJSON()
		return
	}
	var version models.Version
	version.Platforms = vd.Platforms
	ver, err := version.GetVersionInfo()
	if err != nil {
		vc.Data["json"] = RetResource(true, types.SystemDbErr, err, "wallet success")
		vc.ServeJSON()
		return
	}
	data := map[string]interface{}{
		"id": ver.Id,
		"version_num": ver.VersionNum,
		"platforms": ver.Platforms,
		"decribe": ver.Decribe,
		"download_url": ver.DownloadUrl,
		"is_force": ver.IsForce,
	}
	vc.Data["json"] = RetResource(true, types.ReturnSuccess, data, "获取版本信息成功")
	vc.ServeJSON()
	return
}

