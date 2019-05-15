package routers

import (
	"github.com/astaxie/beego"
	"itemDetail/itemDetailInBeego/controllers"
)

func init() {
	//beego.Router("/detail/?:id", &controllers.MainController{}
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.MainController{})

	//当实现了自定义的Git请求方法，请求付不会访问默认方法
	beego.Router("/login", &controllers.MainController{}, "get:ShowLogin;post:HandleLogin")
	beego.Router("/index", &controllers.MainController{}, "get:ShowIndex")
	//添加文章
	beego.Router("/addArticle", &controllers.MainController{}, "get:ShowAdd;post:HandleAdd")
}
