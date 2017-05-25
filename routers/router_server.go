package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/controllers"
	"github.com/feisuweb/fastblog/filters"
)

func init() {
	//服务器
	beego.InsertFilter("/server", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/server", &controllers.ServerHandle{}, "*:GetListServer")

	beego.InsertFilter("/server/add", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/server/add", &controllers.ServerHandle{}, "get:AddServer")

	beego.InsertFilter("/server/add", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/server/add", &controllers.ServerHandle{}, "post:PostAddServer")

	beego.InsertFilter("/server/edit", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/server/edit", &controllers.ServerHandle{}, "get:EditServer")

	beego.InsertFilter("/server/edit", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/server/edit", &controllers.ServerHandle{}, "post:PostEditServer")

	beego.InsertFilter("/server/delete", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/server/delete", &controllers.ServerHandle{}, "*:Delete")
}
