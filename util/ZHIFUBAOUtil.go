package util

import (
	"github.com/smartwalle/alipay"
	"ChickGame/def"
	"fmt"
)
var client = alipay.New(def.ZHIFUBAOAPPID, "2088802940812132", def.ZHIFUBAO_KEY, def.ZHIFUBAOprivateKey, false)


//传递一个map
func ZhiPayTo(p *alipay.AliPayTradeWapPay)( string){
	var html, _ = client.TradeWapPay(*p)
	fmt.Print(html)
	htmlToString:=fmt.Sprintf("%s",html)
	return htmlToString
}







