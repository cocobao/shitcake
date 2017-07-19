package routers

import (
	"github.com/astaxie/beego"
	"github.com/cocobao/shitcake/controller"
)

func init() {
	beego.Router("/", &controller.HomeController{})
	beego.Router("/login", &controller.LogController{}, "get:Login")
	beego.Router("/login", &controller.LogController{}, "post:LoginCommit")
	beego.Router("/upload", &controller.Upload{})

}
