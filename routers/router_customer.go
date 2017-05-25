package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/controllers"
	"github.com/feisuweb/fastblog/filters"
)

func init() {
	//客户管理
	beego.InsertFilter("/customer", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/customer/", &controllers.CustomerController{}, "*:GetListCustomer")

	beego.InsertFilter("/customer/update", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/customer/update", &controllers.CustomerController{}, "get:GetUpdateCustomer")

	beego.InsertFilter("/customer/update", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/customer/update", &controllers.CustomerController{}, "post:PostUpdateCustomer")

	beego.InsertFilter("/customer/add", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/customer/add", &controllers.CustomerController{}, "get:GetAddCustomer")

	beego.InsertFilter("/customer/add", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/customer/add", &controllers.CustomerController{}, "post:PostAddCustomer")

	beego.InsertFilter("/customer/delete", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/customer/delete", &controllers.CustomerController{}, "post:PostDeleteCustomer")

}
