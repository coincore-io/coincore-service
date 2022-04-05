package models

import (
	"coinwallet/common"
	"coinwallet/types/notebook"
	"github.com/astaxie/beego/orm"
	"time"
)

type AddrNoteBook struct {
	Id           int64      `orm:"pk;column(id);auto;size(11)" description:"ID" json:"id"`
	DeviceId     string     `orm:"column(device_id);size(512)" description:"设备ID" json:"device_id"`
	Name         string     `orm:"column(name);size(512)" description:"地址名称" json:"name"`
	AssetName    string     `orm:"column(asset_name);size(512)" description:"币种名称" json:"asset_name"`
	Memo         string     `orm:"column(memo);size(512)" description:"备注" json:"memo"`
	Addr     	 string     `orm:"column(addr);size(512)" description:"地址" json:"addr"`
	IsRemoved    int8       `orm:"column(is_removed);default(0)" description:"是否删除"  json:"is_removed"`
	CreatedAt    time.Time  `orm:"column(created_at);auto_now_add;type(datetime);index" description:"修改时间" json:"created_at"`
	UpdatedAt    time.Time  `orm:"column(updated_at);auto_now_add;type(datetime);index" description:"更新时间" json:"updated_at"`
}

func (this *AddrNoteBook) TableName() string {
	return common.TableName("addr_notebook")
}

func (this *AddrNoteBook) SearchField() []string {
	return []string{"name","asset_name"}
}

func (this *AddrNoteBook) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func (this *AddrNoteBook) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *AddrNoteBook) Insert() (err error, id int64) {
	if id, err = orm.NewOrm().Insert(this); err != nil {
		return err, 0
	}
	return nil, id
}

func (this *AddrNoteBook) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(this)
}

func GetAddrNoteBookById(id int64) *AddrNoteBook {
	var note_book AddrNoteBook
	if err := orm.NewOrm().QueryTable(AddrNoteBook{}).Filter("id", id).RelatedSel().One(&note_book); err != nil {
		return nil
	}
	return &note_book
}

func UpdateAddressInfo(upd_nb_addr notebook.UpdateAddressNoteBookReq) bool {
	var note_book AddrNoteBook
	if err := orm.NewOrm().QueryTable(AddrNoteBook{}).Filter("id", upd_nb_addr.NbId).RelatedSel().One(&note_book); err != nil {
		return false
	}
	if upd_nb_addr.DeviceId != "" {
		note_book.DeviceId = upd_nb_addr.DeviceId
	}
	if upd_nb_addr.Name != "" {
		note_book.Name = upd_nb_addr.Name
	}
	if upd_nb_addr.Name != "" {
		note_book.Name = upd_nb_addr.Name
	}
	if upd_nb_addr.AssetName != "" {
		note_book.AssetName = upd_nb_addr.AssetName
	}
	if upd_nb_addr.Memo != "" {
		note_book.Memo = upd_nb_addr.Memo
	}
	if upd_nb_addr.Address != "" {
		note_book.Addr = upd_nb_addr.Address
	}
	err := note_book.Update()
	if err != nil {
		return false
	}
	return true
}

func GetUserAddressList(device_id string) []*AddrNoteBook {
	var note_addr_list []*AddrNoteBook
	if _, err := orm.NewOrm().QueryTable(AddrNoteBook{}).Filter("device_id", device_id).All(&note_addr_list); err != nil {
		return nil
	}
	return note_addr_list
}


func DelAddress(id int64) bool {
	_, err := orm.NewOrm().QueryTable(AddrNoteBook{}).Filter("id", id).Delete()
	if err != nil {
		return false
	}
	return true
}