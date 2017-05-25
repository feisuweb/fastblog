package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Server struct {
	Id         int64
	Name       string `orm:"size(100)"` //名称
	Ip         string //站点域名
	Port       int64  //端口
	User       string //用户名
	Password   string //密码
	DirName    string
	CreateTime time.Time `orm:"type(datetime);null"` //站点数据创建时间
	UpdateTime time.Time `orm:"type(datetime);null"` //站点数据创建时间
	Status     int64     //状态 1 正常 0 关闭 -1 删除
}

func (m *Server) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Server) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Server) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Server) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Server) GetServerById(id int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", id).One(m)
	return err != orm.ErrNoRows
}
