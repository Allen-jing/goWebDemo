package routers

import (
	"github.com/astaxie/beego"
	"goWebDemo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
