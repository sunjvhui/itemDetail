package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type photo struct {
	Id string
	ImagePath string
}

//返回三项数据， photos, title, price
func detailGET(c *gin.Context)  {
	//模拟处理商品ID
	id := c.Param("id")
	fmt.Printf("id is :%s\n", id)

	//模拟访问数据库
	//初始化Photo
	buffer := []photo{
		photo{
			Id: "pic1",
			ImagePath: "img01.png",
		},
		photo{
			Id: "pic2",
			ImagePath: "img02.png",
		},
		photo{
			Id: "pic3",
			ImagePath: "img03.png",
		},
		photo{
			Id: "pic4",
			ImagePath: "img04.png",
		},
		photo{
			Id: "pic5",
			ImagePath: "img05.png",
		},
		photo{
			Id: "pic6",
			ImagePath: "img06.png",
		},
	}
	//
	//初始化price

	price := 4541100000
	intPrice := price / 100
	decPrice := price % 100
	priceString := fmt.Sprintf("%d.%02d", intPrice, decPrice)

	//返回json
	c.JSON(http.StatusOK, gin.H{
		"photos": buffer,
		"title": "北欧风格实木床",
		"price": priceString,

	})


}

//Gin实现商品详情页
//1,搭建静态HTML服务
//2,实现AJAX服务器部分
//3,实现AJAX HTML部分

func main()  {
	router := gin.Default()

	//设定HTML文件夹路由，用于存放index.html 以及css/js等文件
	router.Static("/html", "./html")

	//路由，设定/detail/:id, 返回json数据
	router.GET("detail/:id", detailGET)


	//服务启动
	router.Run("127.0.0.1:8092")
}