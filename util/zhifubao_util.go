package util

import (
	"github.com/smartwalle/alipay"
	"ChickGame/def"
	"fmt"
	"net/http"
)
var client = alipay.New(def.ZHIFUBAOAPPID, "2088802940812132", def.ZHIFUBAO_KEY, def.ZHIFUBAOprivateKey, false)


//把这个链接返回给前端，前端直接访问这个url
func ZhiPayTohtml(p *alipay.AliPayTradeWapPay)( string){
	var html, _ = client.TradeWapPay(*p)
	fmt.Print(html)
	htmlToString:=fmt.Sprintf("%s",html)
	return htmlToString
}

//支付宝服务器返回，像支付宝服务器返回成功与失败
func ZhiReturnIf(reql *http.Request )string{
	ok, _ := client.VerifySign(reql.Form)
	if ok{
		return "success"
	}else {
		return "false"
	}

}







