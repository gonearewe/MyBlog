package main

import (
	_ "MyBlog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/article/static", "static")
	beego.Run()
}

