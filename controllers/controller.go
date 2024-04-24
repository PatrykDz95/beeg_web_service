package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

var users = map[string]string{
	"1": "user1",
	"2": "user2",
}

func (c *UserController) GetUserByID() {
	id := c.Ctx.Input.Param(":id")
	user, ok := users[id]
	if !ok {
		c.Ctx.Output.SetStatus(404)
		c.Ctx.Output.Body([]byte("User not found"))
		return
	}
	c.Ctx.Output.Body([]byte(user))
}
