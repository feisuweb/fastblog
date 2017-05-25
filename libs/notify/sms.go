package notify

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
	"time"
)

//=====================短信基础=====================
//send sms
func SendSMSNotify(mobile string, content string) error {
	mobile = strings.Trim(mobile, " ")
	if len(mobile) == 0 {
		return nil
	}
	err := SendSMSDuoWei(mobile, content)
	return err
}

//send sms
func SendSMSDuoWei(mobile string, content string) error {
	url := "http://service.winic.org:8009/sys_port/gateway/index.asp"
	contentDataGBK, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(content)), simplifiedchinese.GBK.NewEncoder()))

	//短信自动加签名
	contentGBK := fmt.Sprintf("%s", string(contentDataGBK))
	req := httplib.Post(url)

	req.Param("id", "13926485656")   // 注册的用户名
	req.Param("pwd", "x326342")      // 注册成功后,登录网站使用的密钥
	req.Param("to", mobile)          // 手机号码
	req.Param("content", contentGBK) // 设置短信内容
	req.Param("time", "")            // 为空即时发送，格式：yyyy/mm/dd hh:mm:ss

	req.Header("Content-Type", "application/x-www-form-urlencoded;charset=gbk")

	beego.Info("sms", mobile+" content:"+content)
	req.SetTimeout(10*time.Second, 10*time.Second)
	_, err := req.DoRequest()
	return err
}

//================================客户通知==============================
//发送给客户短信-产品订单通知
func SendToCustomerSMSOrderNotify(m *NotifyInfo) error {
	body := `恭喜，您的订单$OrderNo$已经发货,发货信息:网盘下载地址：$YunpanDownloadUrl$ 提取码：$DownloadCode$ 解压密码：$UnzipPassword$  本地下载地址：$DownloadUrl$`
	body = ReplaceNotifyContent(m, body)
	err := SendSMSNotify(m.UserMobile, body)
	return err
}

//发送给客户邮箱-用户订单通知
func SendToCustomerSMSUserOrderNotify(m *NotifyInfo) error {
	body := `恭喜您,您成功升级为 $ProductName$ ,您的用户账号是 $UserName$ 默认密码:$Password$ 手机号:$UserMobile$ 邮箱 $UserEmail$ ,登录codeshop.com 优惠下载更多商业产品 `
	body = ReplaceNotifyContent(m, body)

	err := SendSMSNotify(m.UserMobile, body)
	return err
}

//================================卖家通知==============================
//发送给卖家短信-产品订单通知
func SendToSellerSMSOrderNotify(m *NotifyInfo) error {

	body := `卖出产品,下单客户:$UserName$，订单号:$OrderNo$ 金额是：$Amount$ 名称：$ProductName$ 客户手机：$UserMobile$，邮箱: $UserEmail$  下单时间:$AddTime$  `

	body = ReplaceNotifyContent(m, body)

	err := SendSMSNotify(m.SellerMobile, body)
	return err
}

//===========================发送给站长用户购买通知=======================
//发送给站长短信-用户订单通知
func SendToMasterSMSUserOrderNotify(m *NotifyInfo) error {

	body := `卖出用户,用户$UserName$购买VIP套餐$ProductName$，订单号:$OrderNo$ 订单金额是：$Amount$ 代理商佣金：$CommissionAmount$  代理商手机:$AgentMobile$  用户手机：$UserMobile$，用户邮箱: $UserEmail$ `

	body = ReplaceNotifyContent(m, body)
	err := SendSMSNotify(MasterMobile, body)
	return err
}

//================================代理商通知==============================
//发送给卖家短信-产品订单通知
func SendToAgentSMSOrderNotify(m *NotifyInfo) error {

	body := `推荐产品$ProductName$被客户$UserName$ 购买,获取佣金:$CommissionAmount$.`

	body = ReplaceNotifyContent(m, body)

	err := SendSMSNotify(m.AgentMobile, body)
	return err
}

//发送给代理商短信-用户订单通知
func SendToAgentSMSUserOrderNotify(m *NotifyInfo) error {
	body := `代理商$UserName$ 在$AddTime$ 购买VIP套餐$ProductName$，获得佣金额是：$CommissionAmount$`
	body = ReplaceNotifyContent(m, body)

	err := SendSMSNotify(m.AgentMobile, body)
	return err
}

//===================密码修改通知==============
//发送给用户的短信-用户密码修改通知
func SendToUserSMSPasswordChangedNotify(m *NotifyInfo) error {
	if m.IsVip { //只有VIP用户才会收到短信通知
		body := `您的账号$UserName$修改了密码.点这里快速冻结账号（24小时有效）: http://www.codeshop.com/user/lock?id=$UserSafeLogId$`
		body = ReplaceNotifyContent(m, body)
		err := SendSMSNotify(m.UserMobile, body)
		return err
	} else {
		return nil
	}

}
