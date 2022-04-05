package news


type News struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	Abstract  string    `json:"abstract"`
	Image     string    `json:"image"`
	Author    string    `json:"author"`
	Views     int64     `json:"views"`
	Likes     int64     `json:"likes"`
	Date      string    `json:"date"`
	Time      string    `json:"time"`
}
