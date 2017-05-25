package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/filters"
	"github.com/feisuweb/fastblog/models"
	"html/template"
)

//全局变量
var (
	agent_id         int64  = 0
	agent_mobile     string = ""
	login_user_id    int64  = 0
	login_user_token string = ""
	login_user_info  models.User
	login_user_type  string = "普通用户"
	web_site_name           = ""
	SiteSettings     map[string]string
	LoginUser        models.User
	IsLogin          bool = false
)

var (
	app_name       = ""
	access_channel = "direct"
	access_client  = "web"
	access_device  = "pc"

	site_weixin_url = ""
	site_www_url    = ""
	site_order_url  = ""
	site_mobile_url = ""
	site_user_url   = ""
	site_image_url  = ""
	site_file_url   = ""
	site_api_url    = ""

	site_pay_url        = ""
	site_pay_scan_url   = ""
	site_pay_notify_url = ""
)

func init() {

	app_name = beego.AppConfig.String("appname")
	//网址
	site_www_url = beego.AppConfig.String("site_www_url")
	site_weixin_url = beego.AppConfig.String("site_weixin_url")
	site_order_url = beego.AppConfig.String("site_order_url")
	site_mobile_url = beego.AppConfig.String("site_mobile_url")
	site_user_url = beego.AppConfig.String("site_user_url")
	site_pay_url = beego.AppConfig.String("site_pay_url")
	site_image_url = beego.AppConfig.String("site_image_url")
	site_file_url = beego.AppConfig.String("site_file_url")
	site_api_url = beego.AppConfig.String("site_api_url")

	//扫码支付URL
	site_pay_scan_url = site_pay_url + "/pay/scan"
	//支付通知URL
	site_pay_notify_url = site_pay_url + "/pay/notify"
	fmt.Println("===========app config============")
	fmt.Println("app name =" + app_name)
	fmt.Println("site www url=" + site_www_url)
	fmt.Println("site weixin url=" + site_weixin_url)
	fmt.Println("site user url=" + site_user_url)
	fmt.Println("site mobile url=" + site_mobile_url)
	fmt.Println("site order url=" + site_order_url)
	fmt.Println("site image url=" + site_image_url)
	fmt.Println("site file url=" + site_file_url)
	fmt.Println("site pay  url=" + site_pay_url)
	fmt.Println("site pay scan  url=" + site_pay_scan_url)
	fmt.Println("site pay notify url=" + site_pay_notify_url)
}

type baseController struct {
	beego.Controller
}

func (this *baseController) Prepare() {
	//判断登录状态
	IsLogin, LoginUser = filters.IsLogin(this.Controller.Ctx)
	this.Data["XsrfToken"] = this.XSRFToken()
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["IsLogin"] = IsLogin
	this.Data["LoginUser"] = LoginUser
	this.Data["CurrentUrl"] = this.Ctx.Request.RequestURI
	this.Data["SiteRootUrl"] = site_www_url
	this.Data["CurrentNavigation"] = this.Ctx.Request.RequestURI

}

func Error(err error) {
	if err != nil {
		panic(err)
		beego.Error(err.Error())
		//os.Exit(1)
	}
}

func (this *baseController) ShowError(errorMsg string) {
	this.Data["ErrorMsg"] = errorMsg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "error.html"
}

func (this *baseController) ShowMsg(msg string) {
	this.Data["Msg"] = msg
	this.Layout = "layout/_site_layout.html"
	this.TplName = "message.html"
}
