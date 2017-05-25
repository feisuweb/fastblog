package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//销售类型模型
type SalePersonType struct {
	Id         int64
	Name       string    `orm:"size(500)"`  //销售类型中文名称
	Desc       string    `orm:"size(1000)"` //销售描述
	ValidTime  int64     //销售类型按天计算有效时间
	Price      float64   //购买价格
	AddTime    time.Time `orm:"auto_now_add;type(datetime)"` //入库时间
	UpdateTime time.Time `orm:"auto_now_add;type(datetime)"`
	Status     int64     //销售类型:0为有效 -1为无效 1为推荐
}

func (m *SalePersonType) GetProductById(productId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", productId).One(m)
	if err != nil {
		return false
	}
	return true
}

func GetSalePersonTypeNameById(salePersonTypeId int64) string {
	info := new(SalePersonType)
	err := info.GetSalePersonTypeById(salePersonTypeId)
	if err {
		return info.Name
	} else {
		return "普通销售"
	}

}
func (m *SalePersonType) GetSalePersonTypeNameById(salePersonTypeId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", salePersonTypeId).One(m)
	if err != nil {
		return false
	}
	return true
}

func (m *SalePersonType) GetSalePersonTypeById(salePersonTypeId int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", salePersonTypeId).One(m)
	if err != nil {
		return false
	}
	return true
}

func (m *SalePersonType) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *SalePersonType) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *SalePersonType) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *SalePersonType) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}
