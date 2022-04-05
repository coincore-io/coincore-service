package form_validate

type MarketAssetForm struct {
	Id        int64      `form:"id"`
	Name      string     `form:"name"`
	Icon      string     `form:"icon"`
	IsCreate  int        `form:"_create"`
}