package main

import (
	_ "ChickGame/routers"
	"github.com/astaxie/beego"
	"ChickGame/def"
	"ChickGame/util"
)
func init(){
	util.InitCache()
}
func main() {
	fmt.a
	fmt.
	beego.SetStaticPath("/MP_verify_oSClQLOUTyzPRg6o.txt","MP_verify_oSClQLOUTyzPRg6o.txt")
	def.Outtradeno=beego.AppConfig.String("outtradeno")
	beego.Run()
}

