package controllers

import (
	//"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	//"strconv"
	//"github.com/dchest/captcha"
	"github.com/feisuweb/fastblog/libs/notify"
	"github.com/feisuweb/fastblog/models"
	"strings"
	"time"
)

///前台页面handle
type SettingHandle struct {
	baseController
}

//用户激活账号
func (this *SettingHandle) Activate() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_activate_resend.html"
}

//用户个人资料
func (this *SettingHandle) GetProfile() {
	var (
		minfo *models.User = new(models.User)
		ret   bool
	)
	ret = minfo.GetUserById(LoginUser.Id)
	if ret {
		this.Data["Profile"] = minfo
	}
	this.Ctx.Output.Header("Cache-Control", "public")

	this.Data["PageTitle"] = "注册用户"
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_profile.html"
}

//修改个人资料-POST
func (this *SettingHandle) PostProfile() {
	this.Ctx.Output.Header("Cache-Control", "public")
	var (
		minfo *models.User = new(models.User)
		err   error
		ret   bool
	)
	nickname := strings.TrimSpace(this.GetString("nickname"))
	email := strings.TrimSpace(this.GetString("email"))
	mobile := strings.TrimSpace(this.GetString("mobile"))
	ret = minfo.GetUserById(LoginUser.Id)
	if ret {
		minfo.Email = email
		minfo.Mobile = mobile
		minfo.Nickname = nickname
		err = minfo.Update("Email", "Mobile", "Nickname")

		if err != nil {
			this.Data["ErrorMsg"] = "修改个人信息失败！"

		} else {
			//重新加载信息
			minfo.GetUserById(LoginUser.Id)
			LoginUser.Mobile = mobile
			LoginUser.Email = email
			LoginUser.Nickname = nickname
			this.Data["Msg"] = "修改个人信息成功！"
		}
		this.Data["Profile"] = minfo
	} else {
		this.Data["ErrorMsg"] = "无法获取个人信息！"
	}
	this.Data["PageTitle"] = "修改个人信息"
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_profile.html"
}

//修改Email
func (this *SettingHandle) GetEmail() {
	var (
		minfo *models.User = new(models.User)
		ret   bool
	)
	ret = minfo.GetUserById(LoginUser.Id)
	if ret {
		this.Data["Profile"] = minfo
	}
	this.Ctx.Output.Header("Cache-Control", "public")

	this.Data["PageTitle"] = "注册用户"
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_email.html"
}

//修改Email-POST
func (this *SettingHandle) PostEmail() {
	this.Ctx.Output.Header("Cache-Control", "public")
	var (
		minfo *models.User = new(models.User)
		err   error
		ret   bool
	)
	nickname := strings.TrimSpace(this.GetString("nickname"))
	email := strings.TrimSpace(this.GetString("email"))
	mobile := strings.TrimSpace(this.GetString("mobile"))
	ret = minfo.GetUserById(LoginUser.Id)
	if ret {
		minfo.Email = email
		minfo.Mobile = mobile
		minfo.Nickname = nickname
		err = minfo.Update("Email", "Mobile", "Nickname")

		if err != nil {
			this.Data["ErrorMsg"] = "修改个人信息失败！"

		} else {
			//重新加载信息
			minfo.GetUserById(LoginUser.Id)
			LoginUser.Mobile = mobile
			LoginUser.Email = email
			LoginUser.Nickname = nickname
			this.Data["Msg"] = "修改个人信息成功！"
		}
		this.Data["Profile"] = minfo
	} else {
		this.Data["ErrorMsg"] = "无法获取个人信息！"
	}
	this.Data["PageTitle"] = "修改个人信息"
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_email.html"
}

//修改Mobile
func (this *SettingHandle) GetMobile() {
	var (
		minfo *models.User = new(models.User)
		ret   bool
	)
	ret = minfo.GetUserById(LoginUser.Id)
	if ret {
		this.Data["Profile"] = minfo
	}
	this.Ctx.Output.Header("Cache-Control", "public")

	this.Data["PageTitle"] = "修改手机号"
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_mobile.html"
}

//修改Mobile-POST
func (this *SettingHandle) PostMobile() {
	this.Ctx.Output.Header("Cache-Control", "public")
	var (
		minfo *models.User = new(models.User)
		err   error
		ret   bool
	)
	mobile := strings.TrimSpace(this.GetString("mobile"))
	ret = minfo.GetUserById(LoginUser.Id)
	if ret {
		minfo.Mobile = mobile
		err = minfo.Update("Mobile")

		if err != nil {
			this.Data["ErrorMsg"] = "修改个人信息失败！"

		} else {
			//重新加载信息
			minfo.GetUserById(LoginUser.Id)
			LoginUser.Mobile = mobile
			this.Data["Msg"] = "修改手机号成功！"
		}
		this.Data["Profile"] = minfo
	} else {
		this.Data["ErrorMsg"] = "无法获取个人信息！"
	}
	this.Data["PageTitle"] = "修改手机号"
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_mobile.html"
}

//修改密码
func (this *SettingHandle) GetSettingPassword() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Data["PageTitle"] = "修改密码"
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_password.html"
}

//修改密码-POST
func (this *SettingHandle) PostSettingPassword() {
	this.Ctx.Output.Header("Cache-Control", "public")

	var (
		minfo *models.User = new(models.User)
		err   bool
		ip    string
	)
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	oldPassword := strings.TrimSpace(this.GetString("oldpassword"))
	newPassword := strings.TrimSpace(this.GetString("newpassword"))
	err = minfo.ChangePassword(LoginUser.Id, oldPassword, newPassword)

	if err {
		//发送密码修改通知给用户
		t := time.Now().Format("2006-01-02 15:04:05")
		var ni notify.NotifyInfo
		ni.UserName = LoginUser.UserName
		ni.UserEmail = LoginUser.Email
		ni.UserMobile = LoginUser.Mobile
		ni.UserWeixinOpenId = LoginUser.WeixinOpenId
		ni.ChangePasswordTime = t
		ni.ChangePasswordIp = ip
		ni.ChangePasswordNewPassword = newPassword

		go notify.SendToUserPasswordChangedNotify(&ni)

		this.SetSecureCookie(
			beego.AppConfig.String("cookie.secure"),
			beego.AppConfig.String("cookie.token"),
			"", -1,
			"/",
			beego.AppConfig.String("cookie.domain"),
			false,
			true)
		this.Redirect("/login", 302)
		return
	} else {

	}

	this.Data["ErrorMsg"] = "修改密码失败！"
	this.Data["PageTitle"] = "修改密码"
	this.Layout = "layout/_user_settings_layout.html"
	this.TplName = "user/settings/_password.html"
}
