package controllers

import (
	"github.com/astaxie/beego"
	"github.com/smartwalle/alipay"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"ChickGame/util"
	"ChickGame/def"
)

//以下皆为测试用例！！！！
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["json"] = "beego.me"
	c.ServeJSON()
}

// 是异步返回的url，加上验证！！！是支付宝服务器直接访问的路由， 返回给支付宝服务器的必须是“success”和其他
func (c *MainController) GetUrl() {
	fmt.Print("aaaaaaaa")
	req1 := c.Ctx.Request
	returnstring := util.ZhiReturnIf(req1)
	c.Ctx.WriteString(returnstring)
}

//支付宝支付 示例
func (c *MainController) ZHIFUBAO() {
	var client = alipay.New(def.ZHIFUBAOAPPID, "132123", def.ZHIFUBAO_KEY, def.ZHIFUBAOprivateKey, false)
	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = "http://www.baidu.com"
	p.Subject = "这是测试"
	p.OutTradeNo = "23423423121wqeqw"
	p.TotalAmount = "10.00"
	p.ProductCode = "商品编码"
	html, _ := client.TradeWapPay(p)
	fmt.Print(html)
	// 将html输出到浏览器
	c.Data["json"] = html
	c.ServeJSON()

}

//微信支付示例
func GetGzpt() {
	userMap := &util.StringMap{}
	(*userMap)["appid"] = def.WEIXINAPPID
	(*userMap)["mch_id"] = def.WEIXINMCH_ID
	(*userMap)["nonce_str"] = util.GetRandomString()
	(*userMap)["body"] = "erewrwe"
	(*userMap)["out_trade_no"] = "123456"
	(*userMap)["total_fee"] = "12"
	(*userMap)["spbill_create_ip"] = "123.12.12.123"
	(*userMap)["trade_type"] = "APP"
	(*userMap)["notify_url"] = "http://www.weixin.qq.com/wxpay/pay.php"
	(*userMap)["sign_type"] = "MD5"
	(*userMap)["openid"] = "openid"
	xml := util.MapToxml(userMap)
	response, _ := http.Post("https://api.mch.weixin.qq.com/sandbox/pay/unifiedorder", "application/xml;charset=utf-8", strings.NewReader(xml))
	defer response.Body.Close()
	token_body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(token_body))
}

//微信支付
func GetH5() {
	userMap := &util.StringMap{}
	(*userMap)["appId"] = def.WEIXINAPPID
	(*userMap)["timeStamp"] = "21312"
	(*userMap)["nonceStr"] = util.GetRandomString()
	(*userMap)["package"] = "erewrwe"
	(*userMap)["sign_type"] = "MD5"
	(*userMap)["paySign"] = "MD5"
	xml := util.MapToxml(userMap)
	response, _ := http.Post("https://api.mch.weixin.qq.com/sandbox/pay/unifiedorder", "application/xml;charset=utf-8", strings.NewReader(xml))
	defer response.Body.Close()
	token_body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(token_body))
}
