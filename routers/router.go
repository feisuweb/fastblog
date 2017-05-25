package routers

import (
	"github.com/astaxie/beego"
	"github.com/dchest/captcha"
	"github.com/feisuweb/fastblog/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	beego.Handler("/captcha/*.png", captcha.Server(240, 80)) //注册验证码服务，验证码图片的宽高为240 x 80
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{}, "get:GetLogin")
	beego.Router("/login", &controllers.UserController{}, "post:PostLogin")
	beego.Router("/register", &controllers.UserController{}, "get:GetRegister")
	beego.Router("/register", &controllers.UserController{}, "post:PostRegister")
	beego.Router("/logout", &controllers.UserController{}, "*:Logout")
}
