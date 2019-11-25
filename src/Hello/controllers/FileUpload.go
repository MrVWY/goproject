package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type FormController struct {
	beego.Controller
}

func (c *FormController) Get()  {
	c.TplName = "fileupload.html"
}

func (c *FormController) Post()  {
	f,h,err := c.GetFile("uploadname")
	if err!=nil{
		log.Fatal("getfile err = ",err)
	}
	defer f.Close()
	c.SaveToFile("uploadname","static/upload"+h.Filename)
}