package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/my_test?charset=utf8", 30, 30)
	//映射数据
	orm.RegisterModel(new(User))
	//生成表
	orm.RunSyncdb("default", false, true)
}
