package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/go-sql-driver/mysql"
)

func  main()  {
	router :=gin.Default()
	//gin.Context是gin中最重要的部分，主要用来传递参数，验证request请求，以及返回json的response。
	router.GET("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "LJH" && password == "WAS4163" {
			c.String(0,"Login Suceessfully")
		} else {
			c.String(1,"Login fail")
		}

	})
	router.Run(":8001")
}

