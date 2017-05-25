package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id           int64
	SiteId       int64
	CategoryId   int64     //分类ID
	Title        string    `orm:"size(100)"`       //标题
	Content      string    `orm:"type(text)";null` //内容
	Author       string    //作者
	Photo        string    //图片地址
	ViewCount    int64     //浏览量
	CommentCount int64     //评论量
	Description  string    `orm:"type(text)";null`     //网站描述 SEO描述
	Keywords     string    `orm:"type(text)";null`     //网站描述 SEO关键词
	IsTop        int64     `orm:"default(0)"`          //是否置顶 1置顶
	IsHot        int64     `orm:"default(0)"`          //是否热门 1热门
	CreateTime   time.Time `orm:"type(datetime);null"` //创建时间
	UpdateTime   time.Time `orm:"type(datetime);null"` //更新时间
	Status       int64     `orm:"default(0)"`          //0 未审核 1审核 2回收站
}

func (m *Article) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Article) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Article) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Article) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Article) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Article) GetServerById(id int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", id).One(m)
	return err != orm.ErrNoRows
}
