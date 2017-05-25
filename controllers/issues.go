package controllers

import (
	//"encoding/json"
	//"fmt"
	//"github.com/astaxie/beego"
	//"strconv"
	"github.com/dchest/captcha"
	//"github.com/feisuweb/fastblog/libs/notify"
	//"github.com/feisuweb/fastblog/models"
	//"strings"
	//"time"
)

///前台页面handle
type IssuesHandle struct {
	baseController
}

//分派给我的工单
func (this *IssuesHandle) GetAssigned() {

	this.Ctx.Output.Header("Cache-Control", "public")
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["PageTitle"] = "注册用户"
	this.Layout = "layout/_user_issues_layout.html"
	this.TplName = "issues/_assigned.html"
}

//我创建的工单
func (this *IssuesHandle) GetCreated() {

	this.Ctx.Output.Header("Cache-Control", "public")
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["PageTitle"] = "注册用户"
	this.Layout = "layout/_user_issues_layout.html"
	this.TplName = "issues/_created.html"
}
