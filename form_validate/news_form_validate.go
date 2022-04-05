package form_validate

type NewsForm struct {
	Id        int64      `form:"id"`
	Title     string     `form:"title"`
	CatName   string     `form:"cat_name"`
	Abstract  string     `form:"abstract"`
	Content   string     `form:"content"`
	Image     string     `form:"image"`
	Author    string     `form:"author"`
	NewsType  int8       `form:"news_type"`  //0:快讯；1:文章
	IsCreate  int        `form:"_create"`
}