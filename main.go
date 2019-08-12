package main

import (
	_ "beegoApi/routers"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		// 加入swagger自动化api文档
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
