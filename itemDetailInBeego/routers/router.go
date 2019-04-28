package routers

import (
	"itemDetail/itemDetailInBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/detail/?:id", &controllers.MainController{})
}
