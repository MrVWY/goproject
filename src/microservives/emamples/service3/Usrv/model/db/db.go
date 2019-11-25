package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"microservives/emamples/service3/Usrv/model"
)

func Initialization()*gorm.DB {
	var err error
	db, err := gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open(fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", Users, Password, Mode, Host, Port, Dbnames))
	if err != nil{
		panic(err)
	}
	return db
}

func C(db *gorm.DB , u *model.User) {
	if err := db.Create(u).Error; err != nil {
		fmt.Println("插入失败", err)
	}
}

func U(db *gorm.DB , u *model.User)  {

}

func R(db *gorm.DB , u *model.User) (error, model.User) {
	z := model.User{}
	isNotFound := db.Where(&model.User{
		Username:   u.Username,
		Password:   u.Password,
	}).First(&z).RecordNotFound()
	if isNotFound {
		return errors.New("The user is not exitis") ,z
	}
	return nil,z
}

func R2(db *gorm.DB , u *model.User) (model.User){
	z := model.User{}
	db.Where(&model.User{
		Username:   u.Username,
		Password:   u.Password,
	}).Find(&z)
	return z
}

func R1(db *gorm.DB , u string) (string, error) {
	var z string
	isNotFound := db.Where("username = ?",u).First(&z).RecordNotFound()
	if isNotFound {
		return z,errors.New("The user is not exitis")
	}
	return z,nil
}

func D(db *gorm.DB , u *model.User)  {
	db.Delete(&u)
}