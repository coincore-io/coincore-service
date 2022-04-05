package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"time"
)

type News struct {
	Id        int64      `orm:"pk;column(id);auto;size(11)" description:"公告ID" json:"id"`
	Title     string     `orm:"column(title);size(256)" description:"公告标题" json:"title"`
	CatName   string     `orm:"column(cat_name);size(256)" description:"类别名称" json:"cat_name"`
	Abstract  string     `orm:"column(abstract);type(text)" description:"公告摘要" json:"abstract"`
	Content   string     `orm:"column(content);type(text)" description:"公告内容" json:"content"`
	Image     string     `orm:"column(image);default(0)" description:"公告封面" json:"image"`
	Author    string     `orm:"column(author);default(blockshop)" description:"公告作者" json:"author"`
	Views     int64      `orm:"column(views);default(0)" description:"公告浏览次数" json:"views"`
	NewsType  int8       `orm:"column(news_type);default(0)" description:"类型"  json:"news_type"`  //0:快讯；1:文章
	IsRemoved int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *News) TableName() string {
	return common.TableName("news")
}

func (this *News) SearchField() []string {
	return []string{"title"}
}

func (this *News) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *News) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *News) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *News) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func (new *News) GetNewsById() (*News, string, error) {
	var news News
	err := news.Query().Filter("id", new.Id).One(&news)
	if err != nil {
		return nil, "获取新闻详情失败", errors.New( "获取新闻详情失败")
	}
	news.Views += 1
	err = news.Update()
	if err != nil {
		return nil, "更新阅读次数失败", errors.New( "更新阅读次数失败")
	}
	return &news, "", nil
}

func GetNewsList(page, page_size int64, n_type int8) ([]*News, int, string) {
	var nl []*News
	filter := orm.NewOrm().QueryTable(&News{}).Filter("news_type", n_type)
	total, err := filter.Count()
	if err != nil {
		return nil, 0, "获取新闻记录数量失败"
	}
	_, err = filter.Limit(page_size, page_size*(page-1)).OrderBy("-id").All(&nl)
	if err != nil {
		return nil, 0, "获取新闻列表失败"
	}
	return nl, int(total), "获取新闻列表成功"
}
