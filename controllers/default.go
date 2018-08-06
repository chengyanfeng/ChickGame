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

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
//这个url 是异步返回的url，加上验证！！！
func(c *MainController) GetUrl(){
	/*req1 :=c.Ctx.Request
	fmt.Print(req1,"--------------req1-------------")
	fmt.Print(req1.Form,"------------req1.Form----------")
	ok, err := client.VerifySign(req1.Form)
	fmt.Println(ok, err)*/
	c.Ctx.WriteString("false")
}

//支付宝支付
func (c *MainController)ZHIFUBAO(){
	var client = alipay.New(def.ZHIFUBAOAPPID, "132123", def.ZHIFUBAO_KEY, def.ZHIFUBAOprivateKey, false)
	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = "http://www.baidu.com"
	p.Subject = "这是测试"
	p.OutTradeNo = "23423423121wqeqw"
	p.TotalAmount = "10.00"
	p.ProductCode = "商品编码"

	var html, _ = client.TradeWapPay(p)
	fmt.Print(html)
	c.Data["html"] = html
	// 将html输出到浏览器

}

//微信支付
func GetGzpt(){
	userMap:=&util.StringMap{}
	(*userMap)["appid"] = def.WEIXINAPPID
	(*userMap)["mch_id"] = def.WEIXINMCH_ID
	(*userMap)["nonce_str"] = util.GetRandomString()
	(*userMap)["body"] = "erewrwe"
	(*userMap)["out_trade_no"] = "123456"
	(*userMap)["total_fee"] = "12"
	(*userMap)["spbill_create_ip"] = "123.12.12.123"
	(*userMap)["trade_type"] = "APP"
	(*userMap)["notify_url"]="http://www.weixin.qq.com/wxpay/pay.php"
	(*userMap)["sign_type"]="MD5"
	xml:=util.MapToxml(userMap)
	response, _ := http.Post("https://api.mch.weixin.qq.com/sandbox/pay/unifiedorder", "application/xml;charset=utf-8", strings.NewReader(xml))
	defer response.Body.Close()
	token_body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(token_body))
}
//微信支付
func GetH5(){
	userMap:=&util.StringMap{}
	(*userMap)["appId"] = def.WEIXINAPPID
	(*userMap)["timeStamp"] = "21312"
	(*userMap)["nonceStr"] = util.GetRandomString()
	(*userMap)["package"] = "erewrwe"
	(*userMap)["sign_type"]="MD5"
	(*userMap)["paySign"]="MD5"
	xml:=util.MapToxml(userMap)
	response, _ := http.Post("https://api.mch.weixin.qq.com/sandbox/pay/unifiedorder", "application/xml;charset=utf-8", strings.NewReader(xml))
	defer response.Body.Close()
	token_body, _ := ioutil.ReadAll(response.Body)
	fmt.Print(string(token_body))
}