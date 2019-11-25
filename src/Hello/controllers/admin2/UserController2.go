package admin2

import (
	"github.com/astaxie/beego"
)

type Usercontroller2 struct {
	beego.Controller
}

func (c *Usercontroller2) Get()  {
	c.Ctx.WriteString("Hello UserController")
}