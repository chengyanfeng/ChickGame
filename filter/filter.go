package filter

import (
	"io/ioutil"
	"github.com/astaxie/beego/context"
	"ChickGame/util"
)

var BaseFilter = func(ctx *context.Context) {
	if ctx.Request.RequestURI == "/return" || ctx.Request.RequestURI == "/login" {
		return
	} else {
		request_body := ctx.Request.Body
		auth_body, _ := ioutil.ReadAll(request_body)
		p := *util.JsonDecode([]byte(string(auth_body)))
		auth := p["auth"]
		if auth == nil {

		}
	}
}
