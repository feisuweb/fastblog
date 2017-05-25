package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id         int64
	SiteId     int64     //站点ID
	Title      string    `orm:"size(100)"`
	DirName    string    `orm:"size(100)"`       // 文件夹名称
	Content    string    `orm:"type(text);null"` //内容
	CreateTime time.Time `orm:"type(datetime);null"`
	UpdateTime time.Time `orm:"type(datetime);null"`
	Status     int64     `orm:"default(2)"` //1开启 2关闭
}

func (m *Category) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Category) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Category) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Category) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
