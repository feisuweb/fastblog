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
}
