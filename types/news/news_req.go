package news

import "coinwallet/types"

type NewsListReq struct {
	types.PageSizeData
	NewsType int8 `json:"news_type"`  // 0:快讯；1:文章
}

type NewsDatailReq struct {
	NewsId int64 `json:"news_id"`
}