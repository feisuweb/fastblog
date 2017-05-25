package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/feisuweb/fastblog/models"
)

func IsLogin(ctx *context.Context) (bool, models.User) {
	token, flag := ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"), beego.AppConfig.String("cookie.token"))
	var userInfo models.User
	if flag {
		flag, userInfo = models.GetUserByToken(token)
	}
	return flag, userInfo
}

var CheckAuthority = func(ctx *context.Context) {
	ok, _ := IsLogin(ctx)
	if !ok {
		ctx.Redirect(302, "/login")
	}
}
