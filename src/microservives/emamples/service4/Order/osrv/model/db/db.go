package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func Initialization()*gorm.DB {
	var err error
	//db, err := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql",fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", Users, Password, Mode, Host, Port, Names))
	if err != nil{
		panic(err)
	}
	return db
}

func R1(db *gorm.DB,username string) ([]Order,error){
	z := []Order{}
	isNotFound := db.Where("username = ?",username).Find(&z).RecordNotFound()
	if isNotFound {
		return z , errors.New("The user is not exitis")
	}
	return z , nil
}

func C1(db *gorm.DB , u *Order ) {
	if err := db.Create(u).Error; err != nil {
		fmt.Println("插入失败", err)
	}
}