package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/controllers"
	"github.com/feisuweb/fastblog/filters"
)

func init() {
	//用户服务
	beego.InsertFilter("/user", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/", &controllers.UserController{}, "*:Index")

	beego.InsertFilter("/user/activate", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/activate", &controllers.UserController{}, "*:Activate")

	beego.Router("/user/buy", &controllers.UserController{}, "*:Buy")

	beego.InsertFilter("/user/upgrade", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/upgrade", &controllers.UserController{}, "*:Upgrade")

	beego.InsertFilter("/user/createvip", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/createvip", &controllers.UserController{}, "*:CreateVip")

	beego.InsertFilter("/user/upgradevip", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/upgradevip", &controllers.UserController{}, "*:UpgradeVip")

}
