package main

import (
	_ "ChickGame/routers"
	"github.com/astaxie/beego"
	"ChickGame/def"
	"ChickGame/util"
	."ChickGame/db"
)
var gorm Gorm
func Init(){
	util.InitCache()
}
func main() {

	def.Outtradeno=beego.AppConfig.String("outtradeno")
	beego.Run()
}
