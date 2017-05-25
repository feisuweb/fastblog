package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/feisuweb/fastblog/libs/utils"
	_ "github.com/go-sql-driver/mysql"
	"path/filepath"
)

func init() {
	//读取配置文件
	configPath := filepath.Join("conf", "database.conf")
	fmt.Println("Config path:" + configPath)
	red, err := utils.GetConfig(configPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	//读取mysql 配置
	mysqlpass := red.Conf["mysql.password"]
	mysqluser := red.Conf["mysql.user"]
	mysqldb := red.Conf["mysql.database"]
	mysqlhost := red.Conf["mysql.host"]
	mysqlport := red.Conf["mysql.port"]
	//密码长度，判断是否已经加密过
	if len(mysqlpass) == 24 {
		mysqlpass, err = utils.Decrypt(mysqlpass)
		if err != nil {
			fmt.Errorf("Decrypt mysql passwd failed.")
			return
		}
	}
	//没有加密密码，则加密一次密码，并写入配置文件
	if len(mysqlpass) != 24 {
		psd, err := utils.Encrypt(mysqlpass)
		if err != nil {
			fmt.Errorf("decrypt passwd failed.%v", psd)
			return
		}
		psd = "\"" + psd + "\""
		red.Set("mysql.password", psd)
	}

	orm.RegisterModelWithPrefix("fastblog_", new(User))
	orm.RegisterModelWithPrefix("fastblog_", new(Article))
	orm.RegisterModelWithPrefix("fastblog_", new(Category))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqluser+":"+mysqlpass+"@tcp("+mysqlhost+":"+mysqlport+")/"+mysqldb+"?charset=utf8&loc=Asia%2FShanghai")

	name := "default" //数据库别名
	force := false    //不强制建数据库
	verbose := true   //打印建表过程
	orm.RunSyncdb(name, force, verbose)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	//管理员
	initAdmin()
}

func initAdmin() {
	//数据初始化
	var (
		flag bool
		user User
	)
	//管理员初始化

	flag = user.GetUserByUserName("admin")
	if !flag {

		//如果没有超高级管理员，则初始化一个。

		user.Id = 1
		user.UserName = "admin"
		user.Password = "123456"
		user.Nickname = "超级管理员"
		user.Avatar = "/static/imgs/avatar.png"
		user.RegisterIp = "127.0.0.1"
		user.Register()
	}
}
