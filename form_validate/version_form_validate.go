package form_validate

type VersionForm struct {
	Id           int64     `form:"id"`
	VersionNum   string    `form:"version_num"`
	Platforms    int64     `form:"platforms"`         // 0: 安卓 1: IOS
	Decribe      string    `form:"decribe"`           // 版本描述
	DownloadUrl  string    `form:"download_url"`      // 下载地址
	IsForce      int64     `form:"is_force"`          // 0: 不强制更新 1: 强制更新
	IsCreate     int       `form:"_create"`
}