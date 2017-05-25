package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"strings"
	"time"
)

type UserOrder struct {
	Id          int64
	OrderNo     string //订单号码
	UserId      int64  //用户ID
	ProductId   int64  //用户类型ID
	ProductName string `orm:"size(1000)"`

	UserName         string `orm:"size(150)"`
	UserMobile       string `orm:"size(150)"`
	UserEmail        string `orm:"size(500)"`
	UserWeixin       string `orm:"size(500)"`
	UserWeixinOpenId string `orm:"size(500)"`

	FromPlatform   string //来自平台 mobile  手机  pc 电脑
	FromChannel    string //订单渠道  direct 直接推荐 email 邮件  weixingroup 微信群 weixincircle 微信朋友圈 qqgroup QQ群  website 网站 bbs 论坛 searchengine 百度搜索
	FromChannelTag string //渠道标记

	AgentId           int64  //代理商ID
	AgentMobile       string //代理商手机号
	AgentEmail        string //代理商邮箱
	AgentName         string //代理商名称
	AgentWeixinOpenId string //代理商微信ID
	AgentWeixin       string //代理商微信号
	RecommendCode     string //推介码

	CommissionAmount float64 //销售佣金
	Price            float64 //用户价格
	Count            int64   //购买数量
	Discount         float64 //折扣金额
	PayAmount        float64 //支付金额  =交易金额 - 抵扣
	Amount           float64
	PayMethod        string    `orm:"size(100)"` //alipay (支付宝支付)  weixin(微信支付) agent （代理商支付）
	PayStatus        int64     // 0  未支付  1 等待支付回调 2 支付成功 3 支付失败
	PayTime          time.Time `orm:"auto_now_add;type(datetime)"`
	AddTime          time.Time `orm:"auto_now_add;type(datetime)"`
	VipExpire        time.Time `orm:"auto_now_add;type(datetime)"` //vip过期时间
	UserType         int64     // 0 普通用户 1 超级用户（终身用户）
	UpdateTime       time.Time `orm:"auto_now_add;type(datetime)"`
	IsSend           int64     // 0 位发货 1  发货  ，这里发货是发送短信和邮件
	Status           int64     // 0  等待支付  1 支付成功 2 已经发货 -1 订单作废
}

func (m *UserOrder) CreateOrder() (int64, bool) {
	var (
		oid int64
	)
	oid, err := orm.NewOrm().Insert(m)
	if err != nil {
		return 0, false
	}
	return oid, true
}

//生成订单号
func (m *UserOrder) GetRandOrderNo() string {

	datetime := time.Now().Format("2006-01-02 15:04:05")
	datetime = strings.Replace(datetime, " ", "", -1)
	datetime = strings.Replace(datetime, ":", "", -1)
	datetime = strings.Replace(datetime, "-", "", -1)

	var l int64
	l = 8
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	rn := string(result)

	no := fmt.Sprintf("%s%sM", datetime, rn)
	return no
}

func (m *UserOrder) GetUserOrderByOrderNo(orderNo string) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("OrderNo", orderNo).One(m)
	if err != nil {
		return false
	}
	return true
}

///最新用户购买列表
func (m *UserOrder) GetLastUserList(pagesize int64) []*UserOrder {
	var info UserOrder
	list := make([]*UserOrder, 0)

	info.Query().OrderBy("-id").Limit(pagesize, 0).All(&list, "Id", "UserName", "UserMobile", "UserType", "AddTime")

	return list
}

//获取等待发货列表
func (m *UserOrder) GetWaitForSendList(pagesize int64) []*UserOrder {
	var info UserOrder
	list := make([]*UserOrder, 0)
	query := info.Query()
	query = query.Filter("status", 1)
	query = query.Filter("is_send", 0).OrderBy("Id")
	query.Limit(pagesize, 0).All(&list, "Id", "ProductId", "PayAmount", "OrderNo", "UserEmail", "UserId", "UserMobile")
	return list
}

func (m *UserOrder) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *UserOrder) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UserOrder) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *UserOrder) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
