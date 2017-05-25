package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type SiteGroup struct {
	Id         int64
	Title      string    `orm:"size(100)"`
	Content    string    `orm:"type(text);null"` //内容
	CreateTime time.Time `orm:"type(datetime);null"`
	UserId     int64
	UpdateTime time.Time `orm:"type(datetime);null"`
	Status     int64     `orm:"default(2)"` //1开启 2关闭
}

func (m *SiteGroup) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *SiteGroup) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *SiteGroup) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *SiteGroup) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
