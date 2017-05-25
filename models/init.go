package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/feisuweb/fastblog/libs/notify"
	"github.com/feisuweb/fastblog/libs/utils"
	_ "github.com/go-sql-driver/mysql"
	"path/filepath"
	"time"
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
	orm.RegisterModelWithPrefix("fastblog_", new(UserOrder))
	orm.RegisterModelWithPrefix("fastblog_", new(UserType))
	orm.RegisterModelWithPrefix("fastblog_", new(PayLog))
	orm.RegisterModelWithPrefix("fastblog_", new(UserSafeLog))

	orm.RegisterModelWithPrefix("fastblog_", new(Issues))
	orm.RegisterModelWithPrefix("fastblog_", new(IssuesLog))

	orm.RegisterModelWithPrefix("fastblog_", new(SiteGroup))
	orm.RegisterModelWithPrefix("fastblog_", new(Site))

	orm.RegisterModelWithPrefix("fastblog_", new(Article))
	orm.RegisterModelWithPrefix("fastblog_", new(Category))

	orm.RegisterModelWithPrefix("fastblog_", new(Server))
	orm.RegisterModelWithPrefix("fastblog_", new(Domain))

	orm.RegisterModelWithPrefix("fastblog_", new(Customer))

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

//升级用户
func UpgradeVip(orderNo string, userId int64, userTypeId int64) bool {
	//升级用户期限
	//如果之前是VIP用户，则进行累加
	//用户期限=已有的期限+购买的期限
	//如果之前不是用户则直接进行更新期限
	var minfo User
	var moinfo UserOrder
	var mtinfo UserType
	var notifyInfo notify.NotifyInfo

	moinfo.GetUserOrderByOrderNo(orderNo)
	minfo.GetUserById(userId)
	mtinfo.GetUserTypeById(userTypeId)
	if moinfo.Status < 1 || moinfo.Status == 2 {
		return false
	}
	//判断是否过期用户
	t := time.Now()
	//判断用户是否过期
	ret := t.Before(minfo.VipExpire)
	if ret {
		//用户过期,在当前基础上增加时间
		t1 := time.Duration(mtinfo.ValidTime) * 24 * time.Hour
		t.Add(t1)
		minfo.VipExpire = t

	} else {
		//用户没有过期，则在这个基础上增加时间
		t2 := time.Duration(mtinfo.ValidTime) * 24 * time.Hour
		minfo.VipExpire.Add(t2)

	}
	//设置用户为VIP
	minfo.IsVip = 1
	minfo.Update()
	//更新用户订单表
	moinfo.IsSend = 1
	moinfo.Status = 2
	moinfo.Update()
	//通知

	//通知信息赋值
	//订单信息
	notifyInfo.OrderNo = orderNo
	notifyInfo.Amount = moinfo.Amount
	notifyInfo.PayMethod = moinfo.PayMethod

	//产品信息
	notifyInfo.ProductId = moinfo.ProductId
	notifyInfo.ProductName = moinfo.ProductName

	//用户信息
	notifyInfo.UserId = userId
	notifyInfo.UserEmail = moinfo.UserEmail
	notifyInfo.UserMobile = moinfo.UserMobile
	notifyInfo.UserName = moinfo.UserName

	//推荐者信息
	notifyInfo.AgentId = moinfo.AgentId
	notifyInfo.AgentName = moinfo.AgentName
	notifyInfo.AgentEmail = moinfo.AgentEmail
	notifyInfo.AgentMobile = moinfo.AgentMobile
	notifyInfo.AgentWeixinOpenId = moinfo.AgentWeixinOpenId

	//给客户发送用户订单通知
	notify.SendToCustomerUserOrderNotify(&notifyInfo)
	//给站长发送用户卖出通知
	notify.SendToMasterUserOrderNotify(&notifyInfo)
	//给推荐者发送用户卖出通知
	notify.SendToAgentUserOrderNotify(&notifyInfo)
	return true
}
