package routers

import (
	"github.com/astaxie/beego"
	"goWebDemo/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "Get:Index")

	beego.Router("/menu", &controllers.MenuController{}, "Get:Index")
	beego.Router("/menu/list", &controllers.MenuController{}, "*:List")
	beego.Router("/menu/edit", &controllers.MenuController{}, "*:Edit")
	beego.Router("/menu/editdo", &controllers.MenuController{}, "*:EditDo")
	beego.Router("/menu/add", &controllers.MenuController{}, "Get:Add")
	beego.Router("/menu/adddo", &controllers.MenuController{}, "*:AddDo")

	//login
	beego.Router("/login", &controllers.LoginController{}, "*:Index")

}
