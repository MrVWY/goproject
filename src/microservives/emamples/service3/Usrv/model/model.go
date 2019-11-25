package model

type User struct {
	Username    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	Createtime int64  `gorm:"column:createtime"`
}
