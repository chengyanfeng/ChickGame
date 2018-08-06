package routers

import (
	"ChickGame/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/return", &controllers.MainController{},"post:GetUrl")
}
