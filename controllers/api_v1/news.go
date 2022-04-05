package api_v1

import (
	db_news "coinwallet/models"
	"coinwallet/types"
	"coinwallet/types/news"
	"encoding/json"
	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

// GetNewsList @Title GetNewsList
// @Description 获取消息列表 PostSendPhoneCode
// @Success 200 status bool, data interface{}, msg string
// @router /get_news_list [post]
func (nc *NewsController) GetNewsList() {
	var nlst news.NewsListReq
	if err := json.Unmarshal(nc.Ctx.Input.RequestBody, &nlst); err != nil {
		nc.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		nc.ServeJSON()
		return
	}
	news_list, total, msg := db_news.GetNewsList(int64(nlst.Page), int64(nlst.PageSize), nlst.NewsType)
	var news_lists []news.News
	for _, value := range news_list {
		news_r := news.News{
			Id:        value.Id,
			Title:     value.Title,
			Abstract:  value.Abstract,
			Image:     value.Image,
			Author:    value.Author,
			Views:     value.Views,
			Date: 	   value.CreatedAt.Format("2006-01-02"),
			Time:  	   value.CreatedAt.Format("15:04:05"),
		}
		news_lists = append(news_lists, news_r)
	}
	data := map[string]interface{}{
		"total":     total,
		"gds_lst":   news_lists,
	}
	nc.Data["json"] = RetResource(true, types.ReturnSuccess, data, msg)
	nc.ServeJSON()
	return
}

// GetNewsDetail @Title GetNewsDetail
// @Description 获取消息详情 GetNewsDetail
// @Success 200 status bool, data interface{}, msg string
// @router /news_detail [post]
func (nc *NewsController) GetNewsDetail() {
	var news_ news.NewsDatailReq
	if err := json.Unmarshal(nc.Ctx.Input.RequestBody, &news_); err != nil {
		nc.Data["json"] = RetResource(false, types.InvalidFormatError, err, "无效的参数格式,请联系客服处理")
		nc.ServeJSON()
		return
	}
	var news_s db_news.News
	news_s.Id = news_.NewsId
	news_ret, msg, err := news_s.GetNewsById()
	if err != nil {
		nc.Data["json"] = RetResource(false, types.SystemDbErr, nil, msg)
		nc.ServeJSON()
		return
	}
	nc.Data["json"] = RetResource(true, types.ReturnSuccess, news_ret, "获取新闻详情成功")
	nc.ServeJSON()
	return
}