package routers

import (
	"ChickGame/controllers"
	"github.com/astaxie/beego"
	."ChickGame/filter"
)

func init() {

	beego.InsertFilter("/*", beego.BeforeRouter,BaseFilter)


	//登陆api
	beego.Router("/login", &controllers.Authentication{},"post:Login")

	//以下为测试用例
	//下面是微信转发
	beego.Router("/main/index", &controllers.MainController{},"get:GetIndex")

	beego.Router("/main/get", &controllers.MainController{},"get:GetTicker")

	beego.Router("/main/get_user_token",&controllers.MainController{},"post:GetToken")


	//下面是微信登陆
	beego.Router("/url", &controllers.MainController{},"get:Url")
	beego.Router("/test", &controllers.MainController{},"get:Get")
	beego.Router("/return", &controllers.MainController{},"post:GetUrl")
}
