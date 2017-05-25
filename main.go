package main

import (
	"github.com/astaxie/beego"
	"github.com/feisuweb/fastblog/models"
	_ "github.com/feisuweb/fastblog/routers"
	"os"
)

func main() {
	//创建附件目录
	os.Mkdir("logs", os.ModePerm)
	os.Mkdir("upload", os.ModePerm)
	os.Mkdir("html", os.ModePerm)
	os.Mkdir("post", os.ModePerm)
	os.Mkdir("upload/images", os.ModePerm)
	os.Mkdir("upload/files", os.ModePerm)
	beego.AddFuncMap("ReplaceMobile", models.ReplaceMobile)
	beego.AddFuncMap("GetUserNameById", models.GetUserNameById)
	beego.SetLogFuncCall(true)
	beego.SetLogger("file", `{"filename":"logs/web.log"}`)
	beego.Info("FastBLog服务已经启动...")
	beego.Run()
}
