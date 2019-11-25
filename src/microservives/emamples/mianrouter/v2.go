package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	V "microservives/emamples/mianrouter/view"
	N "microservives/emamples/mianrouter/middleware/jwt"
	"time"
)

func Auth() gin.HandlerFunc  {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		password := c.Request.FormValue("Password")
		z := &N.User{
			Username: username,
			Password: password,
		}
		d := time.Duration(1000)*time.Second
		token := c.Request.Header.Get("token")
		fmt.Println(token)
		if len(token) == 0{
			c.JSON(http.StatusUnauthorized,gin.H{
				"status":"you hava no token",
			})
			c.Abort()
			return
		}
		_, err := N.ParseToken(token)
		if err != nil{
			token, _ = N.JwtGenerateToken(z,d)
			c.JSON(http.StatusUnauthorized,gin.H{
				"status":"your token is wrong",
				"token":token,
			})
			c.Abort()
		}
		c.Next()
	}
}

func x() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		fmt.Println(username)
		re := regexp.MustCompile("[~!;@#$%^&*(){}|<>\\\\/+\\-【】:\"?'：；‘’“”，。、《》\\]\\[`]")
		isIllegal := re.MatchString(username)
		if isIllegal  {
			c.JSON(http.StatusBadRequest,gin.H{
				"status": "请重新输入",
			})
			c.Abort()// 如果有特殊字符 就终止请求向下走
		}

	}
}
//[~!@#$%^&*(){}|<>\\\\/+\\-【】:\"?'：；‘’“”，。、《》\\]\\[`]
func main() {
	router := gin.Default()

	router.POST("/register",x(), V.S2)
	router.POST("/login",Auth(),x(), V.S3)

	router.POST("/test", x(),func(c *gin.Context) {
		username := c.PostForm("username")
		fmt.Println("you input right username", username)
		c.JSON(http.StatusOK, gin.H{
			"status": "S",
		})
	})

	router.Run(":9000")
}
