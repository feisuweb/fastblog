package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type UserSafeLog struct {
	Id      int64     //id
	Action  string    `orm:"size(100)"` //动作 changepassword，login，logout，buy，download，view
	Content string    `orm:"type(text)"`
	User    *User     `orm:"rel(fk)"` //用户
	Ip      string    `orm:"size(32)"`
	AddTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func (m *UserSafeLog) AddSafeLog(user User, action string, ip string, content string) error {
	var msli UserSafeLog
	msli.User = &user
	msli.Action = action
	msli.Ip = ip
	msli.Content = content
	if _, err := orm.NewOrm().Insert(&msli); err != nil {
		return err
	}
	return nil
}

func (m *UserSafeLog) GetUserSafeLogByUserId(id int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", id).One(m)
	return err != orm.ErrNoRows
}

func (m *UserSafeLog) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *UserSafeLog) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserSafeLog) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *UserSafeLog) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
