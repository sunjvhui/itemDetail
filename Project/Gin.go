package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//RESTful路由GET函数
func helloWordGet(c *gin.Context)  {
	c.String(http.StatusOK, "[gin]Hello, World in GET")
}

//RESTful路由POST函数
func helloWordPost(c *gin.Context)  {
	c.String(http.StatusOK, "[gin]Hello, World in POST")
}

//RESTful路由POST函数
func fetchId(c *gin.Context)  {
	id := c.Param("id")
	c.String(http.StatusOK, fmt.Sprintf("id is :%s\n", id))
}

//组路由
func action1(c *gin.Context)  {
	c.String(http.StatusOK,"action 1")
}
func action2(c *gin.Context)  {
	c.String(http.StatusOK,"action 2")
}
func action3(c *gin.Context)  {
	c.String(http.StatusOK,"action 3")
}

func main()  {
	router := gin.Default()

	//RESTful路由
	router.GET("/RESTful", helloWordGet)
	router.POST("/RESTful", helloWordPost)

	//不支持正则路由
	//提取PATH中的参数
	router.GET("/param/:id", fetchId)

	//组路由
	group1 := router.Group("/g1")
	{
		group1.GET("/action1", action1)
		group1.GET("/action2", action2)
		group1.GET("/action3", action3)
	}

	//启动服务
	router.Run("127.0.0.1:8082")
}