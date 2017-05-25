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

//用户激活账号
func (this *UserController) Activate() {
	this.Ctx.Output.Header("Cache-Control", "public")
	this.Layout = "layout/_user_layout.html"
	this.TplName = "user/_activate_resend.html"
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

//购买用户服务
func (this *UserController) Buy() {

	this.Ctx.Output.Header("Cache-Control", "public")
	var (
		userInfo  models.User
		userOrder models.UserOrder
	)

	userList := userOrder.GetLastUserList(6)

	this.Data["PageTitle"] = "购买用户"
	this.Data["userlist"] = userList

	this.Data["userInfo"] = userInfo

	this.Layout = "layout/_user_layout.html"

	this.TplName = "user/_buy.html"
}

//升级用户服务
func (this *UserController) Upgrade() {

	this.Ctx.Output.Header("Cache-Control", "public")

	var (
		userInfo  models.User
		userOrder models.UserOrder
	)

	userList := userOrder.GetLastUserList(6)

	this.Data["PageTitle"] = "升级用户"
	this.Data["userlist"] = userList

	this.Data["userInfo"] = userInfo

	this.Layout = "layout/_user_layout.html"

	this.TplName = "user/_upgrade.html"
}

//购买VIP用户
func (this *UserController) CreateVip() {

	var (
		info        *models.UserOrder = new(models.UserOrder)
		productInfo *models.UserType  = new(models.UserType)
		minfo       *models.User      = new(models.User)
		payinfo     *models.PayLog    = new(models.PayLog)
		//agentInfo      *models.Agent       = new(models.Agent)
		err          bool
		orderNo      string
		url          string
		user_id      int64
		user_type_id int64
		mobile       string
		email        string
		password     string
		ip           string
	)

	user_id, _ = this.GetInt64("user_id")
	user_type_id, _ = this.GetInt64("vip_type")
	mobile = strings.TrimSpace(this.GetString("mobile"))
	email = strings.TrimSpace(this.GetString("email"))
	password = strings.TrimSpace(this.GetString("password"))

	if !models.ValidateMobile(mobile) {
		this.Abort("手机号格式错误！")
		return
	}
	if !models.ValidateEmail(email) {
		this.Abort("请填写正确格式的邮箱！")
		return
	}
	ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	//检查用户之前是否注册过本网站，注册过，则直接登录
	err = minfo.GetUserByMobileAndEmail(mobile, email)
	if err {
		//如果查询到用户已经存在，则
		user_id = minfo.Id
	} else {
		//注册账号信息
		//默认以邮箱和手机号注册一个用户，用户密码是随机数
		//userName string, password string, mobile string, email string, ip string
		//ipResult := models.TabaoAPI(ip)
		minfo.Email = email
		minfo.Password = password
		minfo.Mobile = mobile
		minfo.Nickname = "用户" + mobile
		minfo.UserName = mobile
		minfo.RegisterIp = ip
		minfo.IsVip = 0
		minfo.IsValidateMobile = 0
		minfo.IsValidateEmail = 0
		minfo.Points = 0
		minfo.Money = 0
		//如果有代理商信息
		//if agent_id > 0 && len(agent_mobile) > 0 {
		//minfo.Id = ""
		//minfo.AgentMobile = ""

		//}
		err := minfo.Register()
		if err {
			user_id = minfo.Id
			//更新代理商数据
			//如果有代理商信息
			//if agent_id > 0 && len(agent_mobile) > 0 {
			//agentInfo.Id = ""
			//agentInfo.AddUserCount()

			//}
		}
	}
	//根据产品ID查询产品信息

	err = productInfo.GetUserTypeById(user_type_id)

	if !err {
		this.Abort("用户类型信息有误，请查验后再提交")
	}

	//如果是VIP用户，则直接判断
	// if minfo.CheckVip(minfo.Id) {
	// 	this.Abort("已经是VIP，无需再次购买！")
	// }
	//判断之前是否已经购买过，购买过则无需再次购买
	//订单创建流程开始
	//获取随机订单号
	orderNo = info.GetRandOrderNo()
	//订单创建
	info.OrderNo = orderNo
	info.ProductId = user_type_id
	info.ProductName = productInfo.Name
	info.UserId = user_id

	info.FromPlatform = "pc"
	info.FromChannel = "direct"
	info.FromChannelTag = "codeshop.com"

	// info.RecommendCode = agentInfo.RecommendCode
	// info.AgentId = agentInfo.Id
	// info.AgentName = agentInfo.AgentName
	// info.AgentWeixinOpenId = agentInfo.WeixinOpenId
	// info.AgentWeixin = agentInfo.Weixin
	// info.AgentEmail = agentInfo.Email
	// info.AgentMobile = agentInfo.Mobile

	info.UserName = minfo.Nickname
	info.UserMobile = minfo.Mobile
	info.UserEmail = minfo.Email
	info.UserWeixin = minfo.Weixin
	info.UserWeixinOpenId = minfo.WeixinOpenId

	info.CommissionAmount = 0
	info.Count = 1
	if minfo.CheckVip(user_id) {
		//VIP 用户，采用用户价购买
		info.Price = productInfo.Price
		info.Discount = 0
	} else {
		//普通用户，采用普通价格购买
		info.Price = productInfo.Price
		info.Discount = 0
	}
	info.PayAmount = info.Price
	info.Amount = info.Price
	info.IsSend = 0
	info.Status = 0
	//创建订单
	orderId, oerr := info.CreateOrder()
	if oerr == false {
		this.Abort("用户订单创建失败")
	}

	//创建微信支付记录
	payinfo.OrderId = orderId
	payinfo.OrderNo = info.OrderNo
	payinfo.PayType = 1 //消费
	payinfo.OrderType = "user"
	payinfo.UserId = user_id
	payinfo.AgentId = info.AgentId
	payinfo.UserName = info.UserName
	payinfo.UserMobile = mobile
	payinfo.UserEmail = email
	payinfo.UserWeixin = info.UserWeixin
	payinfo.Amount = info.Amount
	payinfo.Discount = info.Discount
	payinfo.PayAmount = info.PayAmount
	payinfo.PayMethod = "weixin"
	payinfo.PayBody = "购买用户服务" + info.ProductName + "-优品源码网"
	payinfo.ProductId = info.ProductId
	payinfo.PayStatus = 0
	payinfo.Status = 0
	payinfo.Insert()

	//增加代理商VIP用户数
	//if agent_id > 0 && len(agent_mobile) > 0 {
	//	agentInfo.Id = ""
	//	agentInfo.AddVipUserCount()

	//}

	url = site_pay_scan_url + "?orderno=" + orderNo
	if info.PayAmount > 0 {
		url = site_pay_scan_url + "?orderno=" + orderNo
	} else {
		//直接跳转用户详细页面
		url = fmt.Sprintf("/user/profile/%d")
	}
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")
	mid3 := fmt.Sprintf("%d", user_id)
	this.Ctx.SetCookie("user_id", mid3)

	this.SetSecureCookie(
		beego.AppConfig.String("cookie.secure"),
		beego.AppConfig.String("cookie.token"),
		minfo.Token, 30*24*60*60, "/",
		beego.AppConfig.String("cookie.domain"),
		false,
		true)

	this.Redirect(url, 302)
	return
}

//升级VIP用户
func (this *UserController) UpgradeVip() {

	var (
		info        *models.UserOrder = new(models.UserOrder)
		productInfo *models.UserType  = new(models.UserType)
		minfo       *models.User      = new(models.User)
		payinfo     *models.PayLog    = new(models.PayLog)
		//agentInfo      *models.Agent       = new(models.Agent)
		err          bool
		orderNo      string
		url          string
		user_id      int64
		user_type_id int64
		mobile       string
		email        string
		//ip             string
	)
	user_id = LoginUser.Id
	user_type_id, _ = this.GetInt64("vip_type")
	mobile = LoginUser.Mobile
	email = LoginUser.Email

	//ip = this.Ctx.Request.Header.Get("X-Forwarded-For")
	//检查用户之前是否注册过本网站，注册过，则直接登录
	minfo.GetUserById(user_id)
	//查询用户套餐
	err = productInfo.GetUserTypeById(user_type_id)

	if !err {
		this.Abort("用户类型信息有误，请查验后再提交")
	}

	//如果是VIP用户，则直接判断
	//判断之前是否已经购买过，购买过则无需再次购买
	//订单创建流程开始
	//获取随机订单号
	orderNo = info.GetRandOrderNo()
	//订单创建
	info.OrderNo = orderNo
	info.ProductId = user_type_id
	info.ProductName = productInfo.Name
	info.UserId = user_id

	info.FromPlatform = "pc"
	info.FromChannel = "direct"
	info.FromChannelTag = "codeshop.com"

	// info.RecommendCode = agentInfo.RecommendCode
	// info.AgentId = agentInfo.Id
	// info.AgentName = agentInfo.AgentName
	// info.AgentWeixinOpenId = agentInfo.WeixinOpenId
	// info.AgentWeixin = agentInfo.Weixin
	// info.AgentEmail = agentInfo.Email
	// info.AgentMobile = agentInfo.Mobile

	info.UserName = minfo.Nickname
	info.UserMobile = minfo.Mobile
	info.UserEmail = minfo.Email
	info.UserWeixin = minfo.Weixin
	info.UserWeixinOpenId = minfo.WeixinOpenId

	info.CommissionAmount = 0
	info.Count = 1
	if minfo.CheckVip(user_id) {
		//VIP 用户，采用用户价购买
		info.Price = productInfo.Price
		info.Discount = 0
	} else {
		//普通用户，采用普通价格购买
		info.Price = productInfo.Price
		info.Discount = 0
	}
	info.PayAmount = info.Price
	info.Amount = info.Price
	info.IsSend = 0
	info.Status = 0
	//创建订单
	orderId, oerr := info.CreateOrder()
	if oerr == false {

		this.Abort("用户订单创建失败")
	}

	//创建微信支付记录
	payinfo.OrderId = orderId
	payinfo.OrderNo = info.OrderNo
	payinfo.PayType = 1 //消费
	payinfo.OrderType = "user"
	payinfo.UserId = user_id
	payinfo.AgentId = info.AgentId
	payinfo.UserName = info.UserName
	payinfo.UserMobile = mobile
	payinfo.UserEmail = email
	payinfo.UserWeixin = info.UserWeixin
	payinfo.Amount = info.Amount
	payinfo.Discount = info.Discount
	payinfo.PayAmount = info.PayAmount
	payinfo.PayMethod = "weixin"
	payinfo.PayBody = "购买用户服务" + info.ProductName + "-优品源码网"
	payinfo.ProductId = info.ProductId
	payinfo.PayStatus = 0
	payinfo.Status = 0
	payinfo.Insert()

	//增加代理商VIP用户数
	// if agent_id > 0 && len(agent_mobile) > 0 {
	// 	agentInfo.Id = agent_id
	// 	agentInfo.AddVipUserCount()

	// }

	url = site_pay_scan_url + "?orderno=" + orderNo
	if info.PayAmount > 0 {
		url = site_pay_scan_url + "?orderno=" + orderNo
	} else {
		//直接跳转用户详细页面
		url = fmt.Sprintf("/user/profile/%d")
	}
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")

	this.Redirect(url, 302)
	return
}

//订单支付检查
func (this *UserController) Check() {
	var (
		info  = new(models.UserOrder)
		minfo = new(models.User)
		err   bool
	)
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")

	order_no := strings.TrimSpace(this.GetString("orderno"))
	if order_no == "" {
		this.Abort("404")
		return
	}
	//读取数据
	err = info.GetUserOrderByOrderNo(order_no)
	if err == false || info.Status < 1 {
		this.Abort("404")
		return
	}
	if info.IsSend == 0 && info.Status == 1 {
		//未发货状态,则进行用户增加时间处理
		err = minfo.GetUserById(info.UserId)
		if err == false {
			this.Abort("用户信息不存在，请联系管理员")
			return
		}
		//升级用户
		models.UpgradeVip(info.OrderNo, info.UserId, info.ProductId)
	}
	url := fmt.Sprintf("/user/profile/")
	this.Redirect(url, 302)

}

//前台详细页
func (this *UserController) Profile() {
	var (
		userInfo *models.User = new(models.User)
		is_vip   string
		user_id  int64
	)
	//页面cache控制
	this.Ctx.Output.Header("Cache-Control", "public")
	token, _ := this.Ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
	if IsLogin {

		err2 := userInfo.GetUserByIdAndToken(LoginUser.Id, token)
		if !err2 {
			user_id = 0
		} else {
			//登陆用户，则判断是否为VIP用户
			if userInfo.CheckVip(user_id) {
				is_vip = "VIP用户"

			} else {
				is_vip = "普通用户"
			}
		}
	} else {
		this.Redirect("/login", 302)
		return
	}
	this.Data["user_id"] = LoginUser.Id
	this.Data["userInfo"] = userInfo
	this.Data["is_vip"] = is_vip
	this.Layout = "layout/_user_layout.html"
	this.TplName = "user/_profile.html"
}

//在线充值
func (this *UserController) Recharge() {

	this.Ctx.Output.Header("Cache-Control", "public")
	var (
		userInfo  models.User
		userOrder models.UserOrder
	)

	userList := userOrder.GetLastUserList(6)
	this.Data["PageTitle"] = "在线充值"
	this.Data["userlist"] = userList
	this.Data["userInfo"] = userInfo
	this.Layout = "layout/_user_layout.html"

	this.TplName = "user/_recharge.html"
}
