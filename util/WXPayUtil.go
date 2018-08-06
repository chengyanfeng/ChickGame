package util

import (
	"crypto/md5"
	"sort"
	"fmt"
	"strings"
	"encoding/json"
	"github.com/astaxie/beego"
	"encoding/xml"
	"ChickGame/def"
	"math/rand"
	"time"

	"net/http"
	"io/ioutil"
)

//Map转xml
func MapToxml(userMap *StringMap) string {
	(*userMap)["sign"] = GetSign(userMap)
	buf, _ := xml.Marshal(StringMap(*userMap))
	xml := string(buf)
	xml = strings.Replace(xml, "StringMap", "xml", -1)
	return xml
}

//获取签名
func GetSign(p *StringMap) string {
	sign := ""
	md := md5.New()
	strs := []string{}
	for k := range *p {
		strs = append(strs, k)
	}
	sort.Strings(strs)
	for _, v := range strs {
		sign = sign + v + "=" + (*p)[v] + "&"
	}
	sign = sign + "key=" + def.WEIXINKEY
	fmt.Print(sign)
	md.Write([]byte(sign))
	sign = fmt.Sprintf("%x", md5.Sum([]byte(sign)))
	return strings.ToUpper(sign)

}

// interface 转json
func JsonEncode(v interface{}) (r string) {
	b, err := json.Marshal(v)
	if err != nil {
		Error(err)
	}
	r = string(b)
	return
}

// 记录err信息
func Error(v ...interface{}) {
	beego.Error(v)
}

//string 转P
func JsonDecode(b []byte) (p *map[string]interface{}) {
	p = &map[string]interface{}{}
	err := json.Unmarshal(b, p)
	if err != nil {
		Error("JsonDecode", string(b), err)
	}
	return
}

//生成随机字符串
func GetRandomString() string {
	bytes := []byte(def.WEIXINRANDSTR)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 30; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取token和openid
func GetTokenAndOpenid(code string) (openid string) {

	//获取微信token
	response_token, _ := http.Get("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + def.WEIXINAPPID + "&secret=" + def.WEIXINKEY + "&" + code + "&grant_type=authorization_code")
	//关闭链接
	defer response_token.Body.Close()
	token_body, _ := ioutil.ReadAll(response_token.Body)
	p := *JsonDecode([]byte(string(token_body)))
	refresh_token := p["refresh_token"].(string)
	//直接通过获取的token获取刷新token
	refresh_token_token, _ := http.Get("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=" + def.WEIXINAPPID + "&grant_type=refresh_token&refresh_token=" + refresh_token)
	defer refresh_token_token.Body.Close()
	ticket_body, _ := ioutil.ReadAll(refresh_token_token.Body)
	p = *JsonDecode([]byte(string(ticket_body)))
	/*access_token = p["access_token"].(string)*/
	openid = p["openid"].(string)
	return
}


