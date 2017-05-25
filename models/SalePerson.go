package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/sluu99/uuid"
	"time"
)

type SalePerson struct {
	Id                  int64
	Gender              int64     //性别  0  女 1  男
	Mobile              string    `orm:"unique;size(50)"`  //手机号
	Email               string    `orm:"unique;size(250)"` //邮箱
	Avatar              string    `orm:"size(500)"`        //头像
	SalePersonName      string    `orm:"unique;size(250)"` //销售名称
	Password            string    `orm:"size(32)"`         //密码
	PasswordSalt        string    `orm:"null;size(8)"`
	Nickname            string    `orm:"unique;size(40)"` //昵称
	Pid                 int64     // 推荐销售ID  0  自己注册
	WeixinOpenId        string    `orm:"size(250)"`                   //微信开放平台ID
	Weixin              string    `orm:"size(250)"`                   //微信号码
	City                string    `orm:"size(250)"`                   //代理城市
	Province            string    `orm:"size(250)"`                   //代理城市
	Region              string    `orm:"size(250)"`                   //代理区域
	Address             string    `orm:"size(500)"`                   //代理地址
	LastLoginTime       time.Time `orm:"auto_now_add;type(datetime)"` //最后登录时间
	RegisterTime        time.Time `orm:"auto_now_add;type(datetime)"` //注册时间
	VipExpire           time.Time `orm:"auto_now_add;type(datetime)"` //vip过期时间
	SalePersonType      int64     // 1 置业顾问 2 汽车顾问
	LoginTimes          int64     //登录次数
	LastLoginIp         string    `orm:"size(32)"` //最后登录IP
	RegisterIp          string    `orm:"size(32)"` //第一次注册时候的IP
	IsVip               int64     //是否为VIP
	IsValidateMobile    int64     //是否验证手机号
	IsValidateEmail     int64     //是否邮箱地址
	Points              int64     //销售积分
	Money               int64     //金钱数量
	VipLevel            int64     //VIP等级
	RecommendCode       string    //推荐码
	AgentMobile         string    //代理商手机号
	AgentId             int64     //代理商销售ID
	Token               string    //token
	AddTime             time.Time `orm:"auto_now_add;type(datetime)"`
	SalePersonActivated int64     `orm:"default(0);size(2)"` //1 激活 0 未激活
	Status              int64     // 0  正常  -1 封号  1 限制登录
}

func GetSalePersonByToken(token string) (bool, SalePerson) {
	o := orm.NewOrm()
	var salePersonInfo SalePerson
	err := o.QueryTable(salePersonInfo).Filter("Token", token).One(&salePersonInfo)
	return err != orm.ErrNoRows, salePersonInfo
}

func GetSalePersonNameById(salePersonId int64) string {
	info := new(SalePerson)
	err := info.GetSalePersonById(salePersonId)

	if err {
		return info.SalePersonName
	} else {
		return ""
	}

}

func (m *SalePerson) ChangePassword(id int64, oldPassword string, newPassword string) bool {
	var (
		pwd  string
		pwd2 string
	)
	err := m.GetSalePersonById(id)
	if err {
		salt := m.PasswordSalt

		pwd = Md5(oldPassword + salt)
		if pwd == m.Password { //如果老密码正确，则修改新密码
			pwd2 = Md5(newPassword + salt)
			m.Password = pwd2
			m.Update("password")
			return true
		} else {
			return false
		}

	}
	return false
}

func (m *SalePerson) Register() bool {

	o := orm.NewOrm()
	var pwd string
	var token = uuid.Rand().Hex()
	salt := GetRandomSalt()
	pwd = Md5(m.Password + salt)
	m.PasswordSalt = salt
	m.Password = pwd
	m.Token = token
	m.Avatar = "/static/img/avatar_default.png"
	_, err := o.Insert(m)
	return err != orm.ErrNoRows
}

func (m *SalePerson) Login(salePersonname string, password string, ip string) bool {
	var pwd string
	err := m.GetSalePersonBySalePersonName(salePersonname)
	if err {

		salt := m.PasswordSalt
		pwd = Md5(password + salt)
		if m.Password == pwd {

			if len(m.Token) < 8 {
				var token = uuid.Rand().Hex()
				m.Token = token
				//更新登录信息
				m.LastLoginIp = ip
				m.LastLoginTime = time.Now()
				m.Update("Token", "LastLoginIp", "LastLoginTime")
			} else {
				//更新登录信息
				m.LastLoginIp = ip
				m.LastLoginTime = time.Now()
				m.Update("LastLoginIp", "LastLoginTime")
			}

			//记录登录日志
			return true
		} else {
			return false
		}

	}

	return false
}

func (m *SalePerson) CheckVip(salePersonId int64) bool {
	o := orm.NewOrm()
	var mem SalePerson
	err := o.QueryTable(mem).Filter("Id", salePersonId).Filter("IsVip", 1).One(&mem)
	if err != orm.ErrNoRows {
		return false
	}
	//判断是否过期销售
	t := time.Now()
	//判断销售是否过期
	ret := t.Before(mem.VipExpire)
	if ret {
		return true
	}
	return false

}

func (m *SalePerson) GetSalePersonByIdAndToken(salePersonId int64, token string) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", salePersonId).Filter("Token", token).One(m)
	return err != orm.ErrNoRows

}

func (m *SalePerson) GetSalePersonBySalePersonName(salePersonname string) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("SalePersonName", salePersonname).One(m)
	return err != orm.ErrNoRows
}

func (m *SalePerson) GetSalePersonById(id int64) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Id", id).One(m)
	return err != orm.ErrNoRows
}

func (m *SalePerson) GetSalePersonByMobileAndEmail(mobile string, email string) bool {
	o := orm.NewOrm()
	err := o.QueryTable(m).Filter("Email", email).Filter("Mobile", mobile).One(m)
	return err != orm.ErrNoRows
}

func (m *SalePerson) GetSalePersonByMobileOrEmail(mobile string, email string) bool {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.Or("Email", email).Or("Mobile", mobile)
	qs := o.QueryTable(m)
	qs = qs.SetCond(cond1)
	err := qs.One(m)
	return err != orm.ErrNoRows
}

///最新列表
func (m *SalePerson) GetLastSalePersonList(pagesize int64) []*SalePerson {
	var info SalePerson
	list := make([]*SalePerson, 0)

	info.Query().OrderBy("-id").Limit(pagesize, 0).All(&list, "Id", "SalePersonName", "Mobile", "SalePersonType", "AddTime")

	return list
}

func (m *SalePerson) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *SalePerson) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *SalePerson) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *SalePerson) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
