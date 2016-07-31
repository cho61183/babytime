package routers

import (
	"github.com/astaxie/beego"
	"github.com/cho61183/babytime/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/about", &controllers.AboutController{})
}
