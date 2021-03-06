package models

import (
	"coinwallet/common"
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"strings"
)

type AdminUser struct {
	Id         int    `orm:"column(id);auto;size(11)" description:"表ID" json:"id"`
	Username   string `orm:"column(username);size(30)" description:"用户名" json:"username"`
	Password   string `orm:"column(password);size(255)" description:"密码" json:"password"`
	Nickname   string `orm:"column(nickname);size(30)" description:"昵称" json:"nickname"`
	Avatar     string `orm:"column(avatar);size(255)" description:"头像" json:"avatar"`
	Role       string `orm:"column(role);size(200)" description:"角色" json:"role"`
	Status     int8   `orm:"column(status);size(1)" description:"是否启用 0：否 1：是" json:"status"`
	MerchantId int    `orm:"column(merchant_id);size(11);default(0)" description:"商户ID;0表示管理后台" json:"merchant_id"`
	DeleteTime int    `orm:"column(delete_time);;size(10);default(0)" description:"删除时间" json:"delete_time"`
}

//自定义table 名称
func (*AdminUser) TableName() string {
	return common.TableName("admin_user")
}

//定义模型的可搜索字段
func (*AdminUser) SearchField() []string {
	return []string{"nickname", "username"}
}

//禁止删除的数据id
func (*AdminUser) NoDeletionId() []int {
	return []int{}
}

//定义模型可作为条件的字段
func (*AdminUser) WhereField() []string {
	return []string{}
}

//定义可做为时间范围查询的字段
func (*AdminUser) TimeField() []string {
	return []string{}
}

//获取加密字符串，用在登录的时候加密处理
func (adminUser *AdminUser) GetSignStrByAdminUser(ctx *context.Context) string {
	ua := ctx.Input.Header("user-agent")
	return fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf("%d%s%s", adminUser.Id, adminUser.Username, ua))))
}

//获取已授权url
func (adminUser *AdminUser) GetAuthUrl() map[string]interface{} {
	var (
		urlArr orm.ParamsList
	)
	authUrl := make(map[string]interface{})

	o := orm.NewOrm()
	qs := o.QueryTable(new(AdminRole))

	_, err := qs.Filter("id__in", strings.Split(adminUser.Role, ",")).Filter("status", 1).ValuesFlat(&urlArr, "url")
	if err == nil {
		urlIdStr := ""
		for k, row := range urlArr {
			urlStr, ok := row.(string)
			if ok {
				if k == 0 {
					urlIdStr = urlStr
				} else {
					urlIdStr += "," + urlStr
				}
			}
		}
		urlIdArr := strings.Split(urlIdStr, ",")

		var authUrlArr orm.ParamsList

		if len(urlIdStr) > 0 {
			o = orm.NewOrm()
			qs = o.QueryTable(new(AdminMenu))
			_, err := qs.Filter("id__in", urlIdArr).ValuesFlat(&authUrlArr, "url")
			if err == nil {
				for k, row := range authUrlArr {
					val, ok := row.(string)
					if ok {
						authUrl[val] = k
					}
				}
			}
		}
		return authUrl
	}
	return authUrl
}

//获取当前用户已授权的显示菜单
func (adminUser *AdminUser) GetShowMenu() map[int]orm.Params {
	var maps []orm.Params
	returnMaps := make(map[int]orm.Params)
	o := orm.NewOrm()

	if adminUser.Id == 1 {
		_, err := o.QueryTable(new(AdminMenu)).Filter("is_show", 1).OrderBy("sort_id", "id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
		if err == nil {
			for _, m := range maps {
				returnMaps[int(m["Id"].(int64))] = m
			}
			return returnMaps
		} else {
			return map[int]orm.Params{}
		}
	}

	var list orm.ParamsList
	_, err := o.QueryTable(new(AdminRole)).Filter("id__in", strings.Split(adminUser.Role, ",")).Filter("status", 1).ValuesFlat(&list, "url")
	if err == nil {
		var urlIdArr []string
		for _, m := range list {
			urlIdArr = append(urlIdArr, strings.Split(m.(string), ",")...)
		}
		_, err := o.QueryTable(new(AdminMenu)).Filter("id__in", urlIdArr).Filter("is_show", 1).OrderBy("sort_id", "id").Values(&maps, "id", "parent_id", "name", "url", "icon", "sort_id")
		if err == nil {
			for _, m := range maps {
				returnMaps[int(m["Id"].(int64))] = m
			}
			return returnMaps
		} else {
			return map[int]orm.Params{}
		}
	} else {
		return map[int]orm.Params{}
	}

}

//用户角色名称
func (adminUser *AdminUser) GetRoleText() map[int]*AdminRole {
	roleIdArr := strings.Split(adminUser.Role, ",")
	var adminRole []*AdminRole
	_, err := orm.NewOrm().QueryTable(new(AdminRole)).Filter("id__in", roleIdArr, "id", "name").All(&adminRole)
	if err != nil {
		return nil
	} else {
		adminRoleMap := make(map[int]*AdminRole)
		for _, v := range adminRole {
			adminRoleMap[v.Id] = v
		}
		return adminRoleMap
	}
}

//获取所有用户
func (*AdminUser) GetAdminUser() []*AdminUser {
	var adminUsers []*AdminUser
	_, err := orm.NewOrm().QueryTable(new(AdminUser)).All(&adminUsers)
	if err == nil {
		return adminUsers
	} else {
		return nil
	}
}
