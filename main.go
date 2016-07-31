package main

import (
	"github.com/astaxie/beego"
	_ "github.com/cho61183/babytime/routers"
)

func main() {
	beego.SetStaticPath("/css", "css")
	beego.SetStaticPath("/js", "js")
	beego.Run()
}
