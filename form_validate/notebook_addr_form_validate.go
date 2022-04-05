package form_validate

type NotebookForm struct {
	Id           int64      `form:"id"`
	DeviceId     string     `form:"device_id"`
	Name         string     `form:"name"`
	AssetName    string     `form:"asset_name"`
	Memo         string     `form:"memo"`
	Addr     	 string     `form:"addr"`
	IsCreate     int        `form:"_create"`
}