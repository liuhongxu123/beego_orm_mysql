package main

import (
	_ "mysql/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

