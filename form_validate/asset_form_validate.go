package form_validate

type AssetForm struct {
	Id             int64      `form:"id"`
	Name           string     `form:"name"`
	ChainId        int64       `form:"chain_id`
	Icon           string     `form:"icon"`
	Unit           int64      `form:"unit"`
	Uni           int64      `form:"uni"`
	IsCreate  	   int        `form:"_create"`
}