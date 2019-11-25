package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type RegularController struct {
	beego.Controller
}

func (c *RegularController) Get()  {
	id := c.Ctx.Input.Param(":id")
	fmt.Println("id = ", id)
}