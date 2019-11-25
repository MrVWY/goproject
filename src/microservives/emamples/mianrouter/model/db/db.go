package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	Z "microservives/emamples/mianrouter/model"
)

func Initialization()*gorm.DB {
	var err error
	db, err := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	return db
}

func C(db *gorm.DB , u *Z.User ) {
	if err := db.Create(u).Error; err != nil {
		fmt.Println("插入失败", err)
	}
}

func U(db *gorm.DB , u *Z.User)  {

}

func R(db *gorm.DB , u *Z.User) (error,Z.User) {
	z := Z.User{}
	isNotFound := db.Where("username = ?",u.Username).First(&z).RecordNotFound()
	if isNotFound {
		return errors.New("The user is not exitis") ,z
	}
	return nil,z
}

func D(db *gorm.DB , u *Z.User)  {
	db.Delete(&u)
}