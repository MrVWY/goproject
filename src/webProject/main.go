package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"webProject/V2"
	//"github.com/go-sql-driver/mysql"
)

func  main()  {
	router :=gin.Default()

	//模板
	router.LoadHTMLGlob("templates/*")
	//gin.Context是gin中最重要的部分，主要用来传递参数，验证request请求，以及返回json的response。
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK,"Hello World")
	})
	
	router.GET("user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK,"Hello %s",name)
	})

	router.GET("user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + "is" + action
		c.String(http.StatusOK,message)
	})

	//http://127.0.0.1:8000/welcome?firstname=1&lastname=2
	router.GET("/welcome", func(c *gin.Context) {
		firtname := c.DefaultQuery("firstname","Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK,"Hello %s %s ",firtname,lastname)
	})

	router.POST("/from_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick","anonymous")

		c.JSON(http.StatusOK,gin.H{
			"status" : gin.H{
				"status_code" : http.StatusOK,
				"status" : "ok",
			},
			"message"  : message,
			"nick"  : nick,
		})
	})

	router.PUT("/put", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page","0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id:%s; page: %s; name: %s; message: %s ",id,page,name,message )
		c.JSON(http.StatusOK,gin.H{
			"status_code" : http.StatusOK,
		})
	})

	//文件上传
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK,"upload.html",gin.H{})
	})
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)
		file,header,err := c.Request.FormFile("upload")

		if err != nil{
			c.String(http.StatusBadRequest,"Bad request")
			return
		}
		filename := header.Filename

		fmt.Println(file,err,filename)

		out,err  := os.Create("./picture.png")
		if err != nil{
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out,file)
		if err != nil{
			log.Fatal(err)
		}
	})

	//重定向
	router.GET("/redict/baidu", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"www.baidu.com")
	})

	//分组路由
	v1 := router.Group("/v1")
	{
		v1.GET("/login", getv1)
	}

	v2 := router.Group("/V2")
	{
		v2.GET("/login",V2.Getv2)
	}

	//中间件，略  https://studygolang.com/articles/11819?fr=sidebar

	//异步协程


	router.Run(":8001")
}

func getv1(c *gin.Context)  {
	c.String(http.StatusOK,"v1 Hello")
}

//中间件函数
