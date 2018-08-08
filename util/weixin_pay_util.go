package util

import (
	"crypto/md5"
	"sort"
	"fmt"
	"strings"
	"encoding/json"
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
//获取登陆的token和openid
func GetTokenAndOpenid(code string) (access_token, openid string) {

	//获取微信token
	response_token, _ := http.Get("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + def.WEIXINAPPID + "&secret=" + def.WEIXINKEY + "&code=" + code + "&grant_type=authorization_code")
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
	access_token = p["access_token"].(string)
	openid = p["openid"].(string)

	if checkToken(access_token, openid) {
		return
	} else {
		return "token is error", "openid is error"
	}

}

//验证token和openid是否有效
func checkToken(access_token, openid string) bool {
	checkToken, _ := http.Get("https://api.weixin.qq.com/sns/auth?access_token=" + access_token + "&openid=" + openid)
	defer checkToken.Body.Close()
	checkToken_body, _ := ioutil.ReadAll(checkToken.Body)
	p := *JsonDecode([]byte(string(checkToken_body)))
	errmsg := p["errmsg"].(string)
	if errmsg == "ok" {
		S("pay_token",access_token)
		return true

	} else {
		return false
	}
}

//获取微信登陆用户信息
func GetUserInfo(code string) (p *map[string]interface{}) {
	access_token, openid := GetTokenAndOpenid(code)
	userInfo, _ := http.Get("https://api.weixin.qq.com/sns/userinfo?access_token="+access_token+"&openid="+openid+"&lang=zh_CN")
	defer userInfo.Body.Close()
	userInfo_body, _ := ioutil.ReadAll(userInfo.Body)
	p = JsonDecode([]byte(string(userInfo_body)))
	return
	}


//根据token来获取tick
func GetTicket(token string)string{
	//从token获取微信ticket
	response_ticket, _ := http.Get("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=" + token + "&type=jsapi")
	defer response_ticket.Body.Close()
	ticket_body, _ := ioutil.ReadAll(response_ticket.Body)
	p := *JsonDecode([]byte(string(ticket_body)))
	ticket := p["ticket"].(string)
	S("ticket", ticket, 100*time.Minute)
	fmt.Println("ticket 是从重新拿的")
	return string(ticket)
}
//
func GetForwardToken()(token string)  {
	//获取微信转发token
	response_token, _ := http.Get("https://api.weixin.qq.com/cgi-bin/token?appid=wx53d52d70ccd6439f&secret=dfb513840c45e387cd869af3887e69cb&grant_type=client_credential", )
	defer response_token.Body.Close()
	token_body, _ := ioutil.ReadAll(response_token.Body)
	p := *JsonDecode([]byte(string(token_body)))
	token = p["access_token"].(string);
	S("forword_token",token)
	fmt.Println("这是从转发获取拿的token")
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