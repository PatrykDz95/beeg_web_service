package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"webservice/controllers"
)

func init() {
	web.Router("/user/:id", &controllers.UserController{}, "get:GetUserByID")
	web.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
}
