package notify

import (
	"github.com/astaxie/beego"
	"strings"
)

//send weixin msg
func SendWeixinNotify(weixinOpenId string, content string) error {

	weixinOpenId = strings.Trim(weixinOpenId, " ")
	if len(weixinOpenId) == 0 {
		return nil
	}
	beego.Info(weixinOpenId + "----weixin----" + content)
	return nil
}

//================================客户通知==============================
//发送给客户短信-产品订单通知
func SendToCustomerWeixinOrderNotify(m *NotifyInfo) error {
	body := `恭喜，您的订单$OrderNo$已经发货,发货信息:网盘下载地址：$YunpanDownloadUrl$ 提取码：$DownloadCode$ 解压密码：$UnzipPassword$  本地下载地址：$DownloadUrl$`
	body = ReplaceNotifyContent(m, body)
	err := SendWeixinNotify(m.UserWeixinOpenId, body)
	return err
}

//发送给客户邮箱-用户订单通知
func SendToCustomerWeixinUserOrderNotify(m *NotifyInfo) error {
	body := `恭喜您,您成功升级为 $ProductName$ ,您的用户账号是 $UserName$ 默认密码:$Password$ 手机号:$UserMobile$ 邮箱 $UserEmail$ ,登录codeshop.com 优惠下载更多商业产品 `
	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.UserWeixinOpenId, body)
	return err
}

//================================卖家通知==============================
//发送给卖家短信-产品订单通知
func SendToSellerWeixinOrderNotify(m *NotifyInfo) error {

	body := `卖出产品,下单客户:$UserName$，订单号:$OrderNo$ 金额是：$Amount$ 名称：$ProductName$ 客户手机：$UserMobile$，邮箱: $UserEmail$  下单时间:$AddTime$  `

	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.SellerWeixinOpenId, body)
	return err
}

//===========================发送给站长用户购买通知=======================
//发送给站长短信-用户订单通知
func SendToMasterWeixinUserOrderNotify(m *NotifyInfo) error {

	body := `卖出用户,用户$UserName$购买VIP套餐$ProductName$，订单号:$OrderNo$ 订单金额是：$Amount$ 代理商佣金：$CommissionAmount$  代理商手机:$AgentMobile$  用户手机：$UserMobile$，用户邮箱: $UserEmail$ `

	body = ReplaceNotifyContent(m, body)
	err := SendWeixinNotify(MasterWeixinOpenId, body)
	return err
}

//================================代理商通知==============================
//发送给卖家短信-产品订单通知
func SendToAgentWeixinOrderNotify(m *NotifyInfo) error {

	body := `推荐产品$ProductName$被客户$UserName$ 购买,获取佣金:$CommissionAmount$.`

	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.AgentWeixinOpenId, body)
	return err
}

//发送给代理商短信-用户订单通知
func SendToAgentWeixinUserOrderNotify(m *NotifyInfo) error {
	body := `代理商$UserName$ 在$AddTime$ 购买VIP套餐$ProductName$，获得佣金额是：$CommissionAmount$`
	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.AgentWeixinOpenId, body)
	return err
}

//===================密码修改通知==============
//发送给代理商短信-用户订单通知
func SendToUserWeixinPasswordChangedNotify(m *NotifyInfo) error {
	body := `尊敬的VIP用户，您在$SiteName$的账号 $UserName$ 于$PassowrdChangeTime$ 修改了密码，新密码为$ChangePasswordNewPassword$ 。如果不是您本人修改，请及时通过找回密码找回，或者点击冻结账号链接立刻冻结账号(24小时内有效)。  <a href="http://www.codeshop.com/user/lock?id=$UserSafeLogId$">冻结账号</a>`
	body = ReplaceNotifyContent(m, body)

	err := SendWeixinNotify(m.UserWeixinOpenId, body)
	return err
}
