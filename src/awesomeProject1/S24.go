package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)
//connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", MyUser,Password, Host, Port, Db )
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	fmt.Println("2")
}

type User struct {
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	//创建时间，时间戳
	create_time int64 `gorm:"column:createtime"`
}

//由于grom是使用的orm映射，所以需要定义要操作的表的model，在go中需要定义一个struct， struct的名字就是对应数据库中的表名，
//注意gorm查找struct名对应数据库中的表名的时候会默认把你的struct中的大写字母转换为小写并加上“s”，所以可以加上 db.SingularTable(true)
//让grom转义struct名字的时候不用加上s。
//我是提前在数据库中创建好表的然后再用grom去查询的，也可以用gorm去创建表，我感觉还是直接在数据库上创建，修改表字段的操作方便，grom只用来查询和更新数据。

func main() {
	fmt.Println("1")
	u := User{
		Username:"tizi365",
		Password:"123456",
		create_time:time.Now().Unix(),
	}
	//插入一条用户数据
	//下面代码会自动生成SQL语句：INSERT INTO `users` (`username`,`password`,`createtime`) VALUES ('tizi365','123456','1540824823')
	if err := db.Create(u).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}

	//查询并返回第一条数据
	//定义需要保存数据的struct变量
	//u = User{}
	//isNotFound := db.Where("username = ?","tizi365").First(&u).RecordNotFound()
	//if isNotFound {
	//	fmt.Println("找不到记录")
	//	return
	//}
	////打印查询的数据
	//fmt.Println(u.Username,u.Password)
}