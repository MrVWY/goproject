package admin

import (
	"github.com/astaxie/beego"
)

type Usercontroller struct {
	beego.Controller
}

func (c *Usercontroller) Get()  {
	c.Ctx.WriteString("Hello UserController")
}