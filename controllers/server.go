package controllers

import (
	"github.com/feisuweb/fastblog/models"
	"strings"
	//"time"
)

///前台页面handle
type ServerHandle struct {
	baseController
}

//服务器列表
func (this *ServerHandle) GetListServer() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "服务器列表"
	this.Layout = "layout/_server_layout.html"
	this.TplName = "server/_list.html"
}

//创建服务器
func (this *ServerHandle) AddServer() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "创建服务器"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "server/_create.html"
}

//保存服务器
func (this *ServerHandle) PostAddServer() {
	var (
		serverModel *models.Server = new(models.Server)
		name        string
		ip          string
		port        int64 = 22
		user        string
		password    string
		dirName     string
		status      int64 = 0
	)

	this.Ctx.Output.Header("Cache-Control", "public")
	name = strings.TrimSpace(this.GetString("name"))
	ip = strings.TrimSpace(this.GetString("ip"))
	port, _ = this.GetInt64("port", 0)
	user = strings.TrimSpace(this.GetString("user"))
	password = strings.TrimSpace(this.GetString("password"))
	dirName = strings.TrimSpace(this.GetString("dir_name"))
	status, _ = this.GetInt64("status", 0)

	this.Data["PageTitle"] = "创建服务器"
	if len(name) == 0 {
		this.ShowAddError("请填写名称！")
		return
	}
	if len(ip) == 0 {
		this.ShowAddError("请填写IP！")
		return
	}
	if port == 0 {
		this.ShowAddError("请填写端口！")
		return
	}

	if len(user) == 0 {
		this.ShowAddError("请填写用户名！")
		return
	}

	if len(password) == 0 {
		this.ShowAddError("请填写密码！")
		return
	}

	//信息
	serverModel.Name = name
	serverModel.Ip = ip
	serverModel.Port = port
	serverModel.User = user
	serverModel.Password = password
	serverModel.DirName = dirName
	serverModel.Status = status
	ret := serverModel.Insert()

	if ret == nil {
		//注册成功，跳转到服务器首页
		this.Redirect("/server", 302)
		return
	} else {
		this.ShowAddError("创建服务器失败！")
		return
	}
}

func (this *ServerHandle) ShowAddError(errorMsg string) {
	this.Data["ErrorMsg"] = errorMsg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "server/_create.html"
}

func (this *ServerHandle) ShowEdit(msg string) {

	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "编辑服务器"
	this.Data["Msg"] = msg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "server/_edit.html"
}

func (this *ServerHandle) ShowEditError(msg string) {

	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "编辑服务器"
	this.Data["ErrorMsg"] = msg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "server/_edit.html"
}

//编辑服务器
func (this *ServerHandle) EditServer() {

}

//保存编辑服务器
func (this *ServerHandle) PostEditServer() {
	var (
		serverModel *models.Server = new(models.Server)
		id          int64
		name        string
		ip          string
		user        string
		password    string
		port        int64
		status      int64
	)

	this.Ctx.Output.Header("Cache-Control", "public")
	id, _ = this.GetInt64("id", 0)
	name = strings.TrimSpace(this.GetString("name"))
	user = strings.TrimSpace(this.GetString("user"))
	ip = strings.TrimSpace(this.GetString("ip"))
	port, _ = this.GetInt64("port", 22)
	password = strings.TrimSpace(this.GetString("password"))
	this.Data["PageTitle"] = "编辑服务器"

	if len(name) == 0 {
		this.ShowEditError("请填写名称！")
		return
	}
	if len(ip) == 0 {
		this.ShowEditError("请填写IP！")
		return
	}
	if port == 0 {
		this.ShowEditError("请填写端口！")
		return
	}

	if len(user) == 0 {
		this.ShowEditError("请填写用户名！")
		return
	}

	if len(password) == 0 {
		this.ShowEditError("请填写密码！")
		return
	}
	serverModel.GetServerById(id)
	serverModel.Name = name
	serverModel.Port = port
	serverModel.Password = password
	serverModel.User = user
	serverModel.Ip = ip
	serverModel.Status = status
	ret := serverModel.Update()
	if ret == nil {
		//编辑成功
		this.ShowEdit("编辑成功！")
		return
	} else {
		this.ShowEdit("编辑失败！")
		return
	}
}

//删除
func (this *ServerHandle) Delete() {

}
