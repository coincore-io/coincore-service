package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Version struct {
	Id           int64     `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id" form:"id"`
	VersionNum   string    `orm:"column(version_num)" json:"version_num" form:"version_num"`
	Platforms    int64     `orm:"column(platforms)" json:"platforms" form:"platforms"`              // 0: 安卓 1: IOS
	Decribe      string    `orm:"column(decribe)" json:"decribe" form:"decribe"`                  // 版本描述
	DownloadUrl  string    `orm:"column(download_url)" json:"download_url" form:"download_url"`        // 下载地址
	IsForce      int64     `orm:"column(is_force)" json:"is_force" form:"is_force"`    // 0: 不强制更新 1: 强制更新
	IsRemove     int64     `orm:"column(is_remove)" json:"is_remove"`                   // 0: 删除 1: 不删除
	CreatedAt    time.Time `orm:"column(created_at);auto_now_add;type(datetime);index"`
	UpdatedAt    time.Time `orm:"column(updated_at);auto_now_add;type(datetime);index"`
}

func(ver *Version) SearchField() []string {
	return []string{"version_num","platforms"}
}

func (ver *Version) TableName() string {
	return common.TableName("version")
}

func (ver *Version) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(ver)
}

func (ver *Version) Insert() error {
	if _, err := orm.NewOrm().Insert(ver); err != nil {
		return err
	}
	return nil
}

func (ver *Version) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(ver, fields...); err != nil {
		return err
	}
	return nil
}

//获得某个版本的信息
func (Self *Version) FetchOne() *Version{
	err := orm.NewOrm().Read(Self)
	if err != nil {
		return nil
	}
	return Self
}

//获得版本控制列表
func (Self *Version) FetchRows(page,pageSize int,condition *orm.Condition) ([]*Version,int64,error) {
	var versions  []*Version
	offset := (page - 1) * pageSize
	o := orm.NewOrm().QueryTable(Self.TableName()).SetCond(condition)
	total,err := o.Count()
	if err != nil {
		return nil,0,err
	}
	_,err = o.OrderBy("-id").Limit(pageSize,offset).All(&versions)
	if err != nil {
		return nil,0,err
	}
	return versions,total,nil
}

func (ver *Version) GetVersionInfo() (*Version, error) {
	version := Version{}
	err := orm.NewOrm().QueryTable(ver.TableName()).
		Filter("Platforms", ver.Platforms).
		OrderBy("-id").Limit(1).
		One(&version)
	if err != nil {
		logs.Info(err)
		return nil, err
	}
	return &version, nil
}
