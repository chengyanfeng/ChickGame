package filter

import (
	"strings"
	"net/http"
	"github.com/astaxie/beego/context"
)

/*var BaseFilter = func(ctx *context.Context) {
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
}*/
func TransparentStatic(ctx *context.Context) {
	if strings.Index(ctx.Request.URL.Path, "v1/") >= 0 {
		return
	}
	http.ServeFile(ctx.ResponseWriter, ctx.Request, "static/"+ctx.Request.URL.Path)
}