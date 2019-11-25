package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type AContext struct {
	echo.Context
}

func (c *AContext) Foo()  {
	println("foo")
}

func (c *AContext) Bar()  {
	println("Bar")
}

//中间件
//记录访问量
var totalRequests = 0
func CountRequest(next echo.HandlerFunc) (echo.HandlerFunc) {
	return func(e echo.Context) error {
		totalRequests++
		e.Response().Header().Add("request", fmt.Sprintf("&d",totalRequests))
		return next(e)
	}
}

func A(c echo.Context) error {
	return c.String(http.StatusOK, "A") //c.HTML(http.StatusOK, "<h3>Hello, World!</h3>")
}

func main() {
	e := echo.New()
	e.Use(middleware.Secure())
	//传递空的  XSSProtection, ContentTypeNosniff, XFrameOptions 或 ContentSecurityPolicy 来禁用这项保护。
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "<h3>Hello, World!</h3>") //c.HTML(http.StatusOK, "<h3>Hello, World!</h3>")
	})
	e.Static("/static","js")

	//注册中间件
	e.Use(CountRequest)
	e.GET("/hello", func(c echo.Context) error {
		content := c.(*AContext)
		content.Foo()
		content.Bar()
		return content.String(20,"ok")
	})

	e.GET("/A", A)   //e.URI(A, 1) 将生成 /users/1

	e.GET("/B", func(c echo.Context) error {
		return c.String(http.StatusOK, "B")
	}).Name = "B"  //e.Reverse("foobar", 1234) 就会生成 /users/1234

	e.GET("/C", func(c echo.Context) error {
		return c.String(http.StatusOK, "C")
	}).Name = "C"

	t := &Template{
		templates:template.Must(template.ParseGlob("./static/templates/*.html")),
	}
	e.Renderer = t
	e.GET("/D", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello2","world")  //hello2是标识哪个html
	})

	//g := e.Group("/admin")
	//g.Static("/static/js","s")
	//g.Use(middleware.BasicAuth(func(username, password string) bool {
	//	if username == "joe" && passwor d == "secret" {
	//		return true
	//	}
	//	return false
	//}))
	//c := g.Group("/a")

	e.Logger.Fatal(e.Start(":1323"))
}