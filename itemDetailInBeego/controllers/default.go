package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"itemDetail/itemDetailInBeego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//模拟id
	id := c.Ctx.Input.Param(":id")
	fmt.Printf("id is : %s\n", id)

	//设定模板
	c.TplName = "index.html"

	//绑定数据
	c.Data["Photos"] = models.GetPhotos()
	c.Data["Title"] = models.GetTitle()
	c.Data["Price"] = models.GetPriceString()
}
