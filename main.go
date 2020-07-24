package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "goWebDemo/models"
	_ "goWebDemo/routers"
	_ "goWebDemo/sysinit"
)

func main() {
	logs.SetLevel(beego.LevelDebug)
	logs.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.Debug("开始初始化")
	beego.Run()

}
