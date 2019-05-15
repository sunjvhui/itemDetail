package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id   int
	Name string
	Pwd  string
}

//文章表
type Article struct {
	Id       int       `orm:"pk;auto"`
	Aname    string    //题目
	Atime    time.Time `orm:"auto_now"`        //上传时间
	Acount   int       `orm:"default(0);null"` //阅读量
	Acontent string    //内容
	Aimg     string    //图片名字
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	//设置数据库基本信息
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/my_test?charset=utf8", 30, 30)
	//映射数据
	orm.RegisterModel(new(User), new(Article))
	//生成表
	orm.RunSyncdb("default", false, true)
}
