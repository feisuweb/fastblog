package controllers

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/feisuweb/fastblog/libs/utils"
	"github.com/feisuweb/fastblog/models"
)

type ArticleController struct {
	baseController
}

func (this *ArticleController) GetList() {
	//文章列表首页
	category := new(models.Category)
	categorys := []orm.Params{}
	//查询条件：缀美文章类型，一级
	category.Query().Filter("Pid", 0).OrderBy("-Sort", "CreateTime").Values(&categorys, "Id", "Title")
	for _, cate := range categorys {
		//二级
		son := []orm.Params{}
		category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "CreateTime").Values(&son, "Id", "Title")
		cate["Son"] = son
	}
	this.Data["category"] = &categorys
	this.Layout = "/article/layout.html"
	this.TplName = "/article/list_article.html"

}

func (this *ArticleController) PostList() {

	page, _ := this.GetInt64("page", 1)
	rows, _ := this.GetInt64("rows", 10)
	start := (page - 1) * rows
	article := new(models.Article)
	articles := []orm.Params{}
	//默认状态为关闭
	status, _ := this.GetInt64("status", -1)
	//分类可以为空，列出所有
	cid, _ := this.GetInt64("cid", -1)
	top, _ := this.GetInt64("top", -1)
	q := article.Query().Filter("Type", 0)
	if status != -1 {
		q = q.Filter("Status", status)
	} else {
		q = q.Filter("Status__lt", 2)
	}
	if cid != -1 {
		q = q.Filter("CategoryId", cid)
	}
	if top != -1 {
		q = q.Filter("IsTop", top)
	}
	q.OrderBy("-IsTop", "-View", "CreateTime").Limit(rows, start).Values(&articles)
	for _, p := range articles {

		category := new(models.Category)
		category.Id = (p["Cid"]).(int64)
		//beego.Trace(category.Id)
		err := category.Read("Id")
		if err != nil {
			p["Cid"] = "空"
		} else {
			p["Cid"] = category.Title
		}
	}
	count, _ := q.Count()

	this.Data["json"] = &map[string]interface{}{"total": count, "rows": &articles}
	this.Layout = "layout/_site_layout.html"
	this.TplName = "user/_register.html"

}

func (this *ArticleController) GetAddArticle() {

	//文章列表首页
	category := new(models.Category)
	categorys := []orm.Params{}
	//查询条件：缀美文章类型，一级
	category.Query().Filter("Pid", 0).OrderBy("-Sort", "CreateTime").Values(&categorys, "Id", "Title")
	for _, cate := range categorys {
		//二级
		son := []orm.Params{}
		category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "CreateTime").Values(&son, "Id", "Title")
		cate["Son"] = son
	}
	this.Data["category"] = &categorys
	this.TplName = "/article/add_article.html"

}

func (this *ArticleController) PostArticle() {

	title := this.GetString("title")
	content := this.GetString("content")
	author := this.GetString("author")
	cid, _ := this.GetInt64("cid", 0)
	if cid == 0 {
		this.ShowError("请选择文章分类")
		return
	}
	status, _ := this.GetInt64("status", 0)
	is_top, _ := this.GetInt64("is_top", 0)
	photo := this.GetString("photo")
	article := new(models.Article)
	article.Title = title
	article.Author = author
	article.Content = content
	article.CategoryId = cid
	article.Status = status
	article.IsTop = is_top
	article.Photo = photo
	article.CreateTime = utils.GetTime()
	err := article.Insert()
	if err != nil {
		this.ShowError(err.Error())
	} else {
		this.ShowMsg("增加成功")
	}
}

func (this *ArticleController) GetEditArticle() {
	//user := this.GetSession("userinfo")
	//if user == nil {
	//	this.ShowError("session失效，请重新进入后台首页")
	//}
	id, _ := this.GetInt64("id", 0)
	if id == 0 {
		this.ShowError("文章id参数问题")
		return
	}

	//文章列表首页
	category := new(models.Category)
	categorys := []orm.Params{}
	//查询条件：缀美文章类型，一级
	category.Query().Filter("Pid", 0).OrderBy("-Sort", "CreateTime").Values(&categorys, "Id", "Title")
	for _, cate := range categorys {
		//二级
		son := []orm.Params{}
		category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "CreateTime").Values(&son, "Id", "Title")
		cate["Son"] = son
	}
	this.Data["category"] = &categorys

	if id == 0 {
		this.ShowError("没有id参数")
		this.StopRun()
	}
	//显示更改页面
	thisarticle := new(models.Article)
	thisarticle.Id = id
	err := thisarticle.Read()
	if err != nil {
		this.ShowError("不存在该文章或者数据库出错")
		this.StopRun()
	}
	this.Data["thisarticle"] = thisarticle
	this.Data["PageTitle"] = "编辑文章"
	this.Layout = "layout/_user_layout.html"
	this.TplName = "/article/_edit_article.html"

}

func (this *ArticleController) PostEditArticle() {
	//user := this.GetSession("userinfo")
	//if user == nil {
	//	this.ShowError("session失效，请重新进入后台首页")
	//}
	id, _ := this.GetInt64("id", 0)
	if id == 0 {
		this.ShowError("文章id参数问题")
	}
	small := this.GetString("small")
	if small == "1" {
		status, _ := this.GetInt64("status", 0)
		article := new(models.Article)
		article.Id = id
		article.Status = status
		article.UpdateTime = utils.GetTime()
		err := article.Update("Status", "UpdateTime")
		if err != nil {
			this.ShowError(err.Error())
		} else {
			this.ShowMsg("更改状态成功")
		}
		this.StopRun()
	}
	title := this.GetString("title")
	content := this.GetString("content")
	author := this.GetString("author")
	cid, _ := this.GetInt64("cid", 0)
	if cid == 0 {
		this.ShowError("目录选择有问题")
	}
	status, _ := this.GetInt64("status", 0)
	top, _ := this.GetInt64("top", 0)
	photo := this.GetString("photo")
	article := new(models.Article)
	article.Id = id
	article.Title = title
	article.Author = author
	article.Content = content
	article.CategoryId = cid
	article.Status = status
	article.IsTop = top
	article.UpdateTime = utils.GetTime()
	var err error
	if photo != "" {
		article.Photo = photo
		err = article.Update("Title", "Content", "Author", "IsTop", "CategoryId", "Status", "UpdateTime", "Photo")
	} else {
		err = article.Update("Title", "Content", "Author", "IsTop", "CategoryId", "Status", "UpdateTime")
	}
	if err != nil {
		this.ShowError(err.Error())
	} else {
		this.ShowMsg("修改成功")
	}

}

func (this *ArticleController) DeleteArticle() {
	id, _ := this.GetInt64("id", -1)
	if id != -1 {
		article := new(models.Article)
		article.Id = id
		article.Status = 2
		article.UpdateTime = utils.GetTime()
		err := article.Update("Status", "UpdateTime")
		if err != nil {
			this.ShowError(err.Error())
		} else {
			this.ShowMsg("送到回收站")
		}
	} else {
		this.ShowError("id参数问题")
	}
}
