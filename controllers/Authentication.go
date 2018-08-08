package controllers

import (
	"time"
	"ChickGame/util"
	"io/ioutil"
	"github.com/astaxie/beego"
	"fmt"
	"ChickGame/models"
)

type Authentication struct {
	beego.Controller
	Auth            string
	Token           string
	OpenId          string
	IdentifyingCode string
	Time            time.Time
}

//手机号登陆
func (c *Authentication) Login() {
	p := bodyToMap(c)
	auth := p["auth"]
	c.Auth = auth.(string)
	user:=&models.User{}
	user.City="aaaa"
	models.Db.Create(user)
	c.Time = time.Now()
	util.S("auth", c)
	a := util.S("auth")
	fmt.Print(a)
}

//微信登陆

//退出
func (c *Authentication) Exit() {
	p := bodyToMap(c)
	auth := p["auth"]
	log := util.Del(util.ToString(auth))
	if log == "ok" {
		c.Data["json"] = log
		c.ServeJSON()
	}

}

//获取验证码
func (c *Authentication) sendIdentifyingCode() {
	//等待接口
}

//请求body转P
func bodyToMap(c *Authentication) (p util.P) {
	request_body := c.Ctx.Request.Body
	auth_body, _ := ioutil.ReadAll(request_body)
	p = *util.JsonDecode([]byte(string(auth_body)))
	return
}
