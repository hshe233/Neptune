package routers

import (
	"hshe.top/Neptune/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{}, "post:Login")

	beego.Router("/home", &controllers.HomeController{})

}
