package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Domain struct {
	Id           int64
	CategoryId   int64     //站点类目
	Name         string    `orm:"size(100)"` //名称
	Domain       string    //站点域名
	DomainTypeId int64     //网站类型ID
	Content      string    `orm:"type(text)";null` //网站简介 放在关于我们模块
	Description  string    `orm:"type(text)";null` //网站描述 SEO描述
	Keywords     string    `orm:"type(text)";null` //网站描述 SEO关键词
	ServerId     int64     //所属服务器
	IcpNumber    string    `orm:"type(text);null"` //备案信息
	Logo         string    `orm:"size(500);null"`  //网站logo
	Stat         string    `orm:"type(text);null"` //统计代码
	View         int64     //站点点击数
	CreateTime   time.Time `orm:"type(datetime);null"` //站点数据创建时间
	UpdateTime   time.Time `orm:"type(datetime);null"` //站点数据更新时间
	Status       int64     //状态 1 正常 0 关闭 -1 删除
}

func (m *Domain) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Domain) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Domain) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Domain) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Domain) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Domain) GetDomainById(id int64, fields ...string) error {
	m.Id = id
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}
