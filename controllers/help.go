package controllers

type HelpController struct {
	baseController
}

func (this *HelpController) Index() {

	this.Data["SiteName"] = "飞速博客系统"
	this.Data["PageTitle"] = "帮助"
	this.Data["Keywords"] = "astaxie@gmail.com"
	this.Data["Description"] = "astaxie@gmail.com"
	this.Data["Author"] = "astaxie@gmail.com"
	this.Layout = "layout/_help_layout.html"
	this.TplName = "help/_index.html"
}