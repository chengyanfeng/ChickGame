package datasource

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	. "ChickGame/util"
	"ChickGame/models"
	"fmt"

)

var Db *gorm.DB

type Gorm struct {
}
func (c *Gorm) Init() {
	var err error
	conn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		Conn["username"],
		Conn["password"],
		Conn["host"],
		Conn["port"],
		Conn["name"],
	)
	Db, err = gorm.Open("mysql", conn)
	//禁止用复数
	Db.SingularTable(true)
	if err != nil {
		panic(err.Error())
	}
	if Db.HasTable("dungouset") {

	} else {
		Db.CreateTable(&models.User{})
	}

}
var Conn = P{
	"username": "root",
	"password": "root",
	"host":     "localhost",
	"port":     3306,
	"name":     "test",

}