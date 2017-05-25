package routers

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/controllers"
	"github.com/feisuweb/fastblog/filters"
)

func init() {
	//工单模块
	//被分派给我的
	beego.InsertFilter("/issues", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/issues", &controllers.IssuesHandle{}, "*:GetAssigned")
	//被分派给我的
	beego.InsertFilter("/issues/assigned", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/issues/assigned", &controllers.IssuesHandle{}, "*:GetAssigned")
	//我创建的
	beego.InsertFilter("/issues/created", beego.BeforeRouter, filters.CheckAuthority)
	beego.Router("/issues/created", &controllers.IssuesHandle{}, "*:GetCreated")

}