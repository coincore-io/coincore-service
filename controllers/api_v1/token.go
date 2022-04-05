package api_v1

import (
	"coinwallet/models"
	"coinwallet/types"
	"coinwallet/types/token"
	"coinwallet/types/wallet"
	"encoding/json"
	"github.com/astaxie/beego"
)

type TokenController struct {
	beego.Controller
}


// HotTokenList @Title HotTokenList
// @Description 热门的 toekn 列表 HotTokenList
// @Success 200 status bool, data interface{}, msg string
// @router /hot_token_list [post]
func (this *TokenController) HotTokenList () {
	token_list, _ := models.GetHotTokenList()
	if token_list == nil {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	} else {
		var ts_list []*token.TokenSoarchRep
		for _, token_item := range token_list {
			tks := &token.TokenSoarchRep {
				Id: token_item.Id,
				AssetId: token_item.AssetId,
				TokenName: token_item.TokenName,
				Icon: token_item.Icon,
				TokenSymbol: token_item.TokenSymbol,
				ContractAddr: token_item.ContractAddr,
				Decimal: token_item.Decimal,
			}
			ts_list = append(ts_list, tks)
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, ts_list, "wallet success")
		this.ServeJSON()
		return
	}
}

// SourchAddToken @Title SourchAddToken
// @Description 搜索添加币种 SourchAddToken
// @Success 200 status bool, data interface{}, msg string
// @router /sourch_add_token [post]
func (this *TokenController) SourchAddToken() {
	var token_s token.SoarchTokenReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &token_s); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := token_s.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	token_list, total, _ := models.GetTokenList(token_s.TokeName, int64(token_s.Page), int64(token_s.PageSize))
	if token_list == nil {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	} else {
		var ts_list []*token.TokenSoarchRep
		for _, token_item := range token_list {
			tks := &token.TokenSoarchRep {
				Id: token_item.Id,
				AssetId: token_item.AssetId,
				TokenName: token_item.TokenName,
				Icon: token_item.Icon,
				TokenSymbol: token_item.TokenSymbol,
				ContractAddr: token_item.ContractAddr,
				Decimal: token_item.Decimal,
			}
			ts_list = append(ts_list, tks)
		}
		data := map[string]interface{}{
			"total":     total,
			"gds_lst":   ts_list,
		}
		this.Data["json"] = RetResource(true, types.ReturnSuccess, data, "wallet success")
		this.ServeJSON()
		return
	}
}


// DeleteWalletToken @Title DeleteWalletToken
// @Description 删除钱包的代币 DeleteWalletToken
// @Success 200 status bool, data interface{}, msg string
// @router /delete_wallet_token [post]
func (this *TokenController) DeleteWalletToken() {
	var delete_token wallet.DeleteWalletTokenReq
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &delete_token); err != nil {
		this.Data["json"] = RetResource(false, types.InvalidFormatError, err, "wallet error")
		this.ServeJSON()
		return
	}
	if code, err := delete_token.ParamCheck(); err != nil {
		this.Data["json"] = RetResource(false, code, nil, "wallet error")
		this.ServeJSON()
		return
	}
	ok := models.DeleteWalletToken(delete_token.DeviceId, delete_token.WalletUuid, delete_token.ContractAddr)
	if ok == false {
		this.Data["json"] = RetResource(false, types.HandleError, nil, "wallet error")
		this.ServeJSON()
		return
	}
	this.Data["json"] = RetResource(true, types.ReturnSuccess, nil, "wallet success")
	this.ServeJSON()
	return
}

