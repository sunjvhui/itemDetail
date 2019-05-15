package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"itemDetail/itemDetailInBeego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	////插入
	////有orm对象
	//o := orm.NewOrm()
	////2要插入数据的接口体对象
	//user := models.User{}
	////对结构体赋值
	//user.Name = "小花"
	//user.Pwd = "123456"
	////插入
	//_, err := o.Insert(&user)
	//if err != nil {
	//	fmt.Println("插入失败", err)
	//	return
	//}

	//模拟id
	//id := c.Ctx.Input.Param(":id")

	//查询
	/*int, err := strconv.Atoi(id)

	//查询
	//orm对象
	o := orm.NewOrm()
	//查询的对象
	user := models.User{}
	//指定查询字段值
	user.Id = int
	errl := o.Read(&user)
	if errl != nil {
		fmt.Println("查询失败", err)
		return
	}

	fmt.Println("user", user)
	*/

	/*//更新
	//ORM对象
	o := orm.NewOrm()
	//需要更新的结构体对象
	user := models.User{}
	//查询到需要更新的数据
	user.Id = 1
	err := o.Read(&user)
	//给数据重新赋值
	if err == nil {
		user.Pwd = "123456"
		//更新数据
		_, err = o.Update(&user)
		if err != nil {
			fmt.Println("更新失败", err)
			return
		}else {
			fmt.Println("更新成功")
		}

	}*/

	/*	//删除
		//ORM对象
		o := orm.NewOrm()
		//删除的对象
		user := models.User{}
		//指定删除的数据
		user.Id = 1
		//删除
		_, err := o.Delete(&user)
		if err != nil {
			fmt.Println("删除错误", err)
			return
		}
	*/
	//
	//fmt.Printf("id is : %s\n", id)
	//
	////设定模板
	//c.TplName = "index.html"
	//
	////绑定数据
	//c.Data["Photos"] = models.GetPhotos()
	//c.Data["Title"] = models.GetTitle()
	//c.Data["Price"] = models.GetPriceString()

	c.TplName = "index.html"
}

func (c *MainController) Post() {
	//拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")
	fmt.Println(userName, pwd)
	//对数据进行校验
	if userName == "" || pwd == "" {
		fmt.Println("数据不能为空")
		c.Redirect("/register", 302)
		return
	}

	//插入数据库
	o := orm.NewOrm()
	user := models.User{}
	user.Name = userName
	user.Pwd = pwd
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println("插入失败")
	}
	//返回到登录界面
	c.TplName = "login.html"

	//c.Ctx.WriteString("注册成功")

}

func (c *MainController) ShowLogin() {
	c.TplName = "login.html"
}

//登录业务
func (c *MainController) HandleLogin() {
	//1拿到数据
	userName := c.GetString("userName")
	pwd := c.GetString("pwd")

	fmt.Println("dasdadasd", userName, pwd)

	//判断数据是否合法d

	if userName == "" || pwd == "" {
		fmt.Println("不能为空")
		c.TplName = "login.html"
		return
	}
	//查询账号密码是否正确

	o := orm.NewOrm()
	user := models.User{}

	user.Name = userName
	err := o.Read(&user, "Name")
	if err != nil {
		fmt.Println("查询失败")
		c.TplName = "login.html"
		return
	} else if user.Pwd != pwd {
		fmt.Println("密码错误")
		c.TplName = "login.html"
		return
	}
	//跳转
	c.Redirect("/index", 302)
}

func (c *MainController) ShowIndex() {
	c.TplName = "index.html"
}

//显示添加文章界面

func (c *MainController) ShowAdd() {
	c.TplName = "add.html"
}

//添加的实现
func (c *MainController) HandleAdd() {

	//拿到数据
	artiName := c.GetString("articleName")
	artiContent := c.GetString("content")
	fmt.Println(artiName, artiContent)
	f, h, err := c.GetFile("uploadname")
	defer f.Close()

	if err != nil {
		fmt.Println("上传失败", err)
		return
	} else {
		c.SaveToFile("uploadname", "./ststic/img/"+h.Filename)
	}

	//判断数据合法
	if artiName == "" || artiContent == "" {
		fmt.Println("数据不能为空")
		return
	}
	//插入数据
	o := orm.NewOrm()
	arti := models.Article{}
	arti.Aname = artiName
	arti.Acontent = artiContent
	arti.Aimg = "./ststic/img/" + h.Filename

	_, err = o.Insert(&arti)
	if err != nil {
		fmt.Println("插入失败", err)
		return
	}
	//返回首页
	c.Redirect("/index", 302)
}
