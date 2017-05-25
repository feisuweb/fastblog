package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户类型模型
type UserType struct {
	Id         int64
	Name       string    `orm:"size(500)"`  //用户类型中文名称
	Desc       string    `orm:"size(1000)"` //用户描述
	ValidTime  int64     //用户类型按天计算有效时间
	Price      float64   //购买价格
	AddTime    time.Time `orm:"auto_now_add;type(datetime)"` //入库时间
	UpdateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Status     int64     //用户类型:0为有效 -1为无效 1为推荐
}

func (m *UserType) GetProductById(productId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", productId).One(m)
	if err != nil {
		return false
	}
	return true
}

func GetUserTypeNameById(userTypeId int64) string {
	info := new(UserType)
	err := info.GetUserTypeById(userTypeId)
	if err {
		return info.Name
	} else {
		return "普通用户"
	}

}
func (m *UserType) GetUserTypeNameById(userTypeId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", userTypeId).One(m)
	if err != nil {
		return false
	}
	return true
}

func (m *UserType) GetUserTypeById(userTypeId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", userTypeId).One(m)
	if err != nil {
		return false
	}
	return true
}

func (m *UserType) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *UserType) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *UserType) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserType) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
