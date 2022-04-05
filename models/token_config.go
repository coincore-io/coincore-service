package models

import (
	"coinwallet/common"
	"github.com/astaxie/beego/orm"
	"time"
)

type TokenConfig struct {
	Id           int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	AssetId      int64      `orm:"column(asset_id);size(11)" description:"资产ID" json:"asset_id"`
	ChaindId     int64      `orm:"column(chaind_id);size(11)" description:"链ID" json:"chaind_id"`
	TokenName    string     `orm:"column(token_name);size(512)" description:"合约名称" json:"token_name"`
	Icon         string     `orm:"column(icon);size(512);" description:"图片ICON" json:"icon"`
	TokenSymbol  string     `orm:"column(token_symbol);size(512)" description:"合约代币标志" json:"token_symbol"`
	ContractAddr string     `orm:"column(contract_addr);size(512)" description:"合约地址" json:"contract_addr"`
	Decimal      int64      `orm:"column(decimal);default(18)" description:"合约精度" json:"decimal"`
	IsHot        int8       `orm:"column(is_hot);default(0)" description:"是否删除"  json:"is_hot"` // 0:不是热门资产；1:是热门资产
	IsRemoved    int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt    time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt    time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *TokenConfig) TableName() string {
	return common.TableName("token_config")
}

func (this *TokenConfig) SearchField() []string {
	return []string{"token_name","token_symbol"}
}

func (this *TokenConfig) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *TokenConfig) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *TokenConfig) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *TokenConfig) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}


func GetHotTokenList() ([]*TokenConfig, error) {
	token_list := make([]*TokenConfig, 0)
	_, err := orm.NewOrm().QueryTable(TokenConfig{}).Filter("is_hot", 1).All(&token_list)
	if err != nil {
		return nil, err
	}
	return token_list, nil
}


func GetTokenList(token_name string, page, page_size int64) ([]*TokenConfig, int64, error) {
	token_list := make([]*TokenConfig, 0)
	filter := orm.NewOrm().QueryTable(TokenConfig{}).Filter("token_name__contains", token_name)
	total, err := filter.Count()
	if err != nil {
		return nil, 0, err
	}
	_, err = filter.Limit(page_size, page_size*(page-1)).OrderBy("-id").All(&token_list)
	if err != nil {
		return nil, 0, err
	}
	return token_list, total, nil
}
