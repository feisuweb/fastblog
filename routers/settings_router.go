package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/controllers"
	"github.com/feisuweb/fastblog/filters"
)

func init() {

	//设置模块
	beego.InsertFilter("/user/settings", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings", &controllers.SettingHandle{}, "get:GetProfile")

	beego.InsertFilter("/user/settings/profile", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings/profile", &controllers.SettingHandle{}, "get:GetProfile")

	beego.InsertFilter("/user/settings/profile", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings/profile", &controllers.SettingHandle{}, "post:PostProfile")

	beego.InsertFilter("/user/settings/email", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings/email", &controllers.SettingHandle{}, "get:GetEmail")

	beego.InsertFilter("/user/settings/email", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings/email", &controllers.SettingHandle{}, "post:PostEmail")

	beego.InsertFilter("/user/settings/email", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings/email", &controllers.SettingHandle{}, "get:GetMobile")

	beego.InsertFilter("/user/settings/mobile", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings/mobile", &controllers.SettingHandle{}, "post:PostMobile")

	beego.InsertFilter("/user/settings/password", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings/password", &controllers.SettingHandle{}, "get:GetSettingPassword")

	beego.InsertFilter("/user/settings/password", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/user/settings/password", &controllers.SettingHandle{}, "post:PostSettingPassword")

}
