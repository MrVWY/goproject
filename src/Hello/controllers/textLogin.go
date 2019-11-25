package controllers

import (
	"github.com/astaxie/beego"
)

type Session2Controller struct {
	beego.Controller
}

type UserInfo struct {
	Username string
	UserPwd string
}

func (c *Session2Controller) Get() {
	Username := c.Ctx.GetCookie("name")
	pwd := c.Ctx.GetCookie("pwd")

	if Username != ""{
		c.Ctx.WriteString("Username" + Username + "pwd" + pwd)
	}else {
		c.TplName = "login.html"
	}
}

//has error
func (c *Session2Controller) Post(){
	u := UserInfo{}
	if err := c.ParseForm(&u); err != nil{
		c.Ctx.WriteString("err")
	}
	c.Ctx.SetCookie("Username",u.Username)
	c.Ctx.SetCookie("Userpwd",u.UserPwd)
	c.SetSession("Username",u.Username)
	c.SetSession("Userpwd",u.UserPwd)
	c.Ctx.WriteString("Username:" + u.Username + " Password:" + u.UserPwd)
}