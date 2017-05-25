package controllers

import (
	//"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//"strconv"
	"github.com/dchest/captcha"
	"github.com/feisuweb/fastblog/models"
	"strings"
	//"time"
)

///前台页面handle
type UserController struct {
	baseController
}

//用户首页
func (this *UserController) Index() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "用户首页"
	this.Layout = "layout/_user_layout.html"
	this.TplName = "user/_index.html"
}

//用户注册
func (this *UserController) GetRegister() {

	this.Ctx.Output.Header("Cache-Control", "public")
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["PageTitle"] = "注册用户"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "user/_register.html"
}

func (this *UserController) ShowRegisterError(errorMsg string) {
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["ErrorMsg"] = errorMsg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "user/_register.html"
}

//用户登录
func (this *UserController) PostRegister() {
	var (
		minfo    *models.User = new(models.User)
		err      bool
		mobile   string
		email    string
		password string
		ip       string
	)

	this.Ctx.Output.Header("Cache-Control", "public")
	id, value := this.GetString("captcha_id"), this.GetString("captcha")
	b := captcha.VerifyString(id, value) //验证码校验
	if !b {
		this.ShowRegisterError("验证码错误！")
		return
	}
	mobile = strings.TrimSpace(this.GetString("mobile"))
	email = strings.TrimSpace(this.GetString("email"))
	password = strings.TrimSpace(this.GetString("password"))
	this.Data["PageTitle"] = "用户登录"
	if !models.ValidateMobile(mobile) {
		this.ShowRegisterError("手机号格式错误！")
		return
	}
	if !models.ValidateEmail(email) {
		this.ShowRegisterError("请填写正确格式的邮箱！")
		return
	}
	if len(password) == 0 {
		this.ShowRegisterError("请填写密码！")
		return
	}
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	//检查用户之前是否注册过本网站，注册过，则直接登录
	err = minfo.GetUserByMobileOrEmail(mobile, email)
	if err {
		//如果查询到用户已经存在，则提示用户已经存在了。
		this.ShowRegisterError("手机号或者邮箱已经注册过用户账号。")
		return
	}

	//注册账号信息
	minfo.Email = email
	minfo.Password = password
	minfo.Mobile = mobile
	minfo.Nickname = mobile
	minfo.UserName = mobile
	minfo.RegisterIp = ip
	minfo.IsVip = 0
	minfo.IsValidateMobile = 0
	minfo.IsValidateEmail = 0
	minfo.Points = 0
	minfo.Money = 0
	ret := minfo.Register()

	if ret {
		//注册成功，跳转到用户首页
		this.Redirect("/user", 302)
		return
	} else {
		this.ShowRegisterError("账号注册失败！")
		return
	}
}
func (this *UserController) ShowLoginError(errorMsg string) {
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["ErrorMsg"] = errorMsg
	this.Data["PageTitle"] = "用户登录 - 登录错误"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "user/_login.html"
}

//用户登录
func (this *UserController) GetLogin() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "用户登录"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "user/_login.html"
}

//用户登录
func (this *UserController) PostLogin() {

	var (
		minfo    *models.User = new(models.User)
		err      bool
		userName string
		password string
		ip       string
	)
	userName = strings.TrimSpace(this.GetString("account"))
	password = strings.TrimSpace(this.GetString("password"))
	this.Data["PageTitle"] = "用户登录"
	if len(userName) == 0 {
		this.ShowLoginError("请填写用户名！")
		return
	}
	if len(password) == 0 {
		this.ShowLoginError("请填写密码！")
		return
	}

	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	err = minfo.Login(userName, password, ip)
	if err {
		//登录成功
		this.SetSecureCookie(
			beego.AppConfig.String("cookie.secure"),
			beego.AppConfig.String("cookie.token"),
			minfo.Token, 30*24*60*60, "/",
			beego.AppConfig.String("cookie.domain"),
			false,
			true)
		mid2 := fmt.Sprintf("%d", minfo.Id)
		this.Ctx.SetCookie("user_id", mid2)
		this.Redirect("/user/", 302)
		return
	} else {
		//登录失败
		this.ShowLoginError("账号或者密码错误")
		return
	}
}

//用户退出登录
func (this *UserController) Logout() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.SetSecureCookie(
		beego.AppConfig.String("cookie.secure"),
		beego.AppConfig.String("cookie.token"),
		"", -1,
		"/",
		beego.AppConfig.String("cookie.domain"),
		false,
		true)
	this.Redirect("/login", 302)
}

//用户找回密码
func (this *UserController) FindPassword() {

	this.Ctx.Output.Header("Cache-Control", "public")

	this.Layout = "layout/_user_layout.html"

	this.TplName = "user/_findpassword.html"
}
