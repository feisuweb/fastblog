package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/controllers"
	"github.com/feisuweb/fastblog/filters"
)

func init() {
	//销售管理
	beego.InsertFilter("/saleperson", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/saleperson/", &controllers.SalePersonController{}, "*:GetListSalePerson")

	beego.InsertFilter("/saleperson/update", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/saleperson/update", &controllers.SalePersonController{}, "get:GetUpdateSalePerson")

	beego.InsertFilter("/saleperson/update", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/saleperson/update", &controllers.SalePersonController{}, "post:PostUpdateSalePerson")

	beego.InsertFilter("/saleperson/add", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/saleperson/add", &controllers.SalePersonController{}, "get:GetAddSalePerson")

	beego.InsertFilter("/saleperson/add", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/saleperson/add", &controllers.SalePersonController{}, "post:PostAddSalePerson")

	beego.InsertFilter("/saleperson/delete", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/saleperson/delete", &controllers.SalePersonController{}, "post:PostDeleteSalePerson")

}
