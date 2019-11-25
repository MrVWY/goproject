package controllers

import (
	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}
type  User struct {
	Id int           `form:"-"`
	Username string `form:"username"`
	UserPwd string          `form:"UserPwd"`
}

func (c *SessionController) Get()  {
	Username := c.GetSession("name")
	Userpwd := c.GetSession("pwd")
	if nameString , ok := Username.(string); ok && nameString != ""{
		c.Ctx.WriteString("Name:"+Username.(string) + "Pwd:"+Userpwd.(string))
	}
	c.TplName = "login.html"
}

func (c *SessionController) Post()  {
	u := User{}
	if err := c.ParseForm(&u); err != nil{

	}
	c.Ctx.WriteString("Username:" + u.Username + "password:" + u.UserPwd)
}
