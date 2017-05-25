package controllers

import (
	//"encoding/json"
	//"fmt"
	//"github.com/astaxie/beego"
	//"strconv"
	"github.com/dchest/captcha"
	"github.com/feisuweb/fastblog/models"
	"strings"
	//"time"
)

///前台页面handle
type CustomerController struct {
	baseController
}

//客户首页
func (this *CustomerController) GetListCustomer() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "客户管理"
	this.Layout = "layout/_customer_layout.html"
	this.TplName = "customer/_list.html"
}

//客户添加
func (this *CustomerController) GetAddCustomer() {

	this.Ctx.Output.Header("Cache-Control", "public")
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["PageTitle"] = "添加客户"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "customer/_add.html"
}

func (this *CustomerController) ShowAddCustomerError(errorMsg string) {
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["ErrorMsg"] = errorMsg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "customer/_add.html"
}

//客户登录
func (this *CustomerController) PostAddCustomer() {
	var (
		minfo    *models.Customer = new(models.Customer)
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
		this.ShowAddCustomerError("验证码错误！")
		return
	}
	mobile = strings.TrimSpace(this.GetString("mobile"))
	email = strings.TrimSpace(this.GetString("email"))
	password = strings.TrimSpace(this.GetString("password"))
	this.Data["PageTitle"] = "客户登录"
	if !models.ValidateMobile(mobile) {
		this.ShowAddCustomerError("手机号格式错误！")
		return
	}
	if !models.ValidateEmail(email) {
		this.ShowAddCustomerError("请填写正确格式的邮箱！")
		return
	}
	if len(password) == 0 {
		this.ShowAddCustomerError("请填写密码！")
		return
	}
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	//检查客户之前是否添加过本网站，添加过，则直接登录
	err = minfo.GetCustomerByMobileOrEmail(mobile, email)
	if err {
		//如果查询到客户已经存在，则提示客户已经存在了。
		this.ShowAddCustomerError("手机号或者邮箱已经添加过客户账号。")
		return
	}

	//添加账号信息
	minfo.Email = email
	minfo.Password = password
	minfo.Mobile = mobile
	minfo.Nickname = mobile
	minfo.CustomerName = mobile
	minfo.AddIp = ip
	minfo.IsValidateMobile = 0
	minfo.IsValidateEmail = 0
	minfo.Points = 0
	minfo.Money = 0
	ret := minfo.AddCustomer()

	if ret {
		//添加成功，跳转到客户首页
		this.Redirect("/customer", 302)
		return
	} else {
		this.ShowAddCustomerError("账号添加失败！")
		return
	}
}

//客户更新
func (this *CustomerController) GetUpdateCustomer() {

	this.Ctx.Output.Header("Cache-Control", "public")
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["PageTitle"] = "更新客户"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "customer/_update.html"
}

func (this *CustomerController) ShowUpdateCustomerError(errorMsg string) {
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["ErrorMsg"] = errorMsg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "customer/_update.html"
}

//客户登录
func (this *CustomerController) PostUpdateCustomer() {
	var (
		minfo    *models.Customer = new(models.Customer)
		err      bool
		mobile   string
		email    string
		password string
	)

	this.Ctx.Output.Header("Cache-Control", "public")
	id, value := this.GetString("captcha_id"), this.GetString("captcha")
	b := captcha.VerifyString(id, value) //验证码校验
	if !b {
		this.ShowUpdateCustomerError("验证码错误！")
		return
	}
	mobile = strings.TrimSpace(this.GetString("mobile"))
	email = strings.TrimSpace(this.GetString("email"))
	password = strings.TrimSpace(this.GetString("password"))
	this.Data["PageTitle"] = "客户登录"
	if !models.ValidateMobile(mobile) {
		this.ShowUpdateCustomerError("手机号格式错误！")
		return
	}
	if !models.ValidateEmail(email) {
		this.ShowUpdateCustomerError("请填写正确格式的邮箱！")
		return
	}
	if len(password) == 0 {
		this.ShowUpdateCustomerError("请填写密码！")
		return
	}
	//检查客户之前是否更新过本网站，更新过，则直接登录
	err = minfo.GetCustomerByMobileOrEmail(mobile, email)
	if err {
		//如果查询到客户已经存在，则提示客户已经存在了。
		this.ShowUpdateCustomerError("手机号或者邮箱已经更新过客户账号。")
		return
	}

	//更新账号信息
	minfo.Email = email
	minfo.Password = password
	minfo.Mobile = mobile
	minfo.Nickname = mobile
	minfo.CustomerName = mobile
	minfo.IsValidateMobile = 0
	minfo.IsValidateEmail = 0
	minfo.Points = 0
	minfo.Money = 0
	ret := minfo.Update()

	if ret != nil {
		//更新成功，跳转到客户首页
		this.Redirect("/customer", 302)
		return
	} else {
		this.ShowUpdateCustomerError("账号更新失败！")
		return
	}
}

//删除客户
func (this *CustomerController) PostDeleteCustomer() {

	this.Ctx.Output.Header("Cache-Control", "public")
	captchaId := captcha.NewLen(6) //验证码长度为6
	this.Data["CaptchaId"] = captchaId
	this.Data["PageTitle"] = "添加客户"
	this.Layout = "layout/_site_layout.html"
	this.TplName = "customer/_add.html"
}
