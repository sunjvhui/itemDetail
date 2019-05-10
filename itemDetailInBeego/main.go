package main

import (
	"github.com/astaxie/beego"
	_ "itemDetail/itemDetailInBeego/models"
	_ "itemDetail/itemDetailInBeego/routers"
)

func main() {
	beego.Run()
}
