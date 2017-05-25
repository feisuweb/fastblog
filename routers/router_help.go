package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/controllers"
)

func init() {
	beego.Router("/help", &controllers.HelpController{}, "*:Index")
}
