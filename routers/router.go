package routers

import (
	"ChickGame/controllers"
	."ChickGame/filter"
	"github.com/astaxie/beego"
)

func init() {

	beego.InsertFilter("/*", beego.BeforeRouter,BaseFilter)


	//登陆api
	beego.Router("/login", &controllers.Authentication{},"post:Login")

	//以下为测试用例
	beego.Router("/", &controllers.MainController{})
	beego.Router("/return", &controllers.MainController{},"post:GetUrl")
}
