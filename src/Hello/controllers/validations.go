package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/validation"
)

type  User1 struct {
	Username string `form:"username"`
	UserPwd string          `form:"UserPwd"`
}

type L struct {
	beego.Controller
}

func (c *L) Post() {
	u := User1{}
	fmt.Println("u=",u)
}

func (c *L) Get()  {
	c.TplName = "login.html"
}