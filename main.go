package main

import (
	_ "ChickGame/routers"
	"github.com/astaxie/beego"
	"ChickGame/def"
)

func main() {
	def.Outtradeno=beego.AppConfig.String("outtradeno")
	beego.Run()
}
