package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Init() *gorm.DB {
	//db, err := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql",fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",Users, Password, Mode, Host, Port, Names))
	if err != nil{
		panic(err)
	}
	return db
}

func C1(db *gorm.DB, orderStatus OrderStatus)  {

}