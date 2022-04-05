package form_validate

type ChainForm struct {
	Id             int64      `form:"id"`
	Name           string     `form:"name"`
	Mark           string     `form:"mark"`
	Icon           string     `form:"icon"`
	IsCreate       int        `form:"_create"`
}