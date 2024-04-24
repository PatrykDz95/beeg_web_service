package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "webservice/routers"
)

func main() {
	web.Run()
}
