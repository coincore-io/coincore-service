package api_v1

import (
	"coinwallet/models"
	"coinwallet/types"
	"coinwallet/types/notebook"
	"encoding/json"
	"github.com/astaxie/beego"
)

type NodeBookController struct {
	beego.Controller
}


// AddNoteBook @Title AddNoteBook
// @Description 添加地址到地址本 AddNoteBook
// @Success 200 status bool, data interface{}, msg string
// @router /add_note_book [post]
func (this *NodeBookController) AddNoteBook() {
	var add_nb_addr notebook.AddAddressNoteBookReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &add_nb_addr); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	ad_n := models.AddrNoteBook{
		DeviceId: add_nb_addr.DeviceId,
		Name: add_nb_addr.Name,
		AssetName: add_nb_addr.AssetName,
		Memo: add_nb_addr.Memo,
		Addr: add_nb_addr.Address,
	}
	err, _ := ad_n.Insert()
	if err != nil {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "wallet success")
	this.ServeJSON()
	return
}


// AddNoteBook @Title UpdNoteBook
// @Description 修改地址信息 UpdNoteBook
// @Success 200 status bool, data interface{}, msg string
// @router /upd_note_book [post]
func (this *NodeBookController) UpdNoteBook() {
	var upd_nb_addr notebook.UpdateAddressNoteBookReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &upd_nb_addr); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	ok := models.UpdateAddressInfo(upd_nb_addr)
	if ok {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "wallet success")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
	this.ServeJSON()
	return
}


// DelNoteBook @Title DelNoteBook
// @Description 删除地址信息 DelNoteBook
// @Success 200 status bool, data interface{}, msg string
// @router /del_note_book [post]
func (this *NodeBookController) DelNoteBook() {
	var del_nb_addr notebook.DelAddressNoteBookReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &del_nb_addr); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	ok := models.DelAddress(del_nb_addr.NbId)
	if ok {
		this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "wallet success")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
	this.ServeJSON()
	return
}


// GetNoteBook @Title GetNoteBook
// @Description 删除地址信息 GetNoteBook
// @Success 200 status bool, data interface{}, msg string
// @router /get_note_book [post]
func (this *NodeBookController) GetNoteBook() {
	var qnb_addr notebook.QueryAddressNoteBookReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &qnb_addr); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	data_list := models.GetUserAddressList(qnb_addr.DeviceId)
	if data_list == nil {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, data_list, "wallet success")
	this.ServeJSON()
	return
}
