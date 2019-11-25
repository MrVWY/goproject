package main

import (
	"fmt"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"regexp"
	V "microservives/emamples/mianrouter/view"

)

func M() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Request.FormValue("username")
		re := regexp.MustCompile("[~!;@#$%^&*(){}|<>\\\\/+\\-【】:\"?'：；‘’“”，。、《》\\]\\[`]")
		isIllegal := re.MatchString(username)
		if isIllegal  {
			c.JSON(http.StatusBadRequest,gin.H{
				"status": "请重新输入",
			})
			c.Abort()// 如果有特殊字符 就终止请求
			return
		}

	}
}
//[~!@#$%^&*(){}|<>\\\\/+\\-【】:\"?'：；‘’“”，。、《》\\]\\[`]
func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret")) //密钥对
	router.Use(sessions.Sessions("mysession",store)) //让router使用name为mysession 密钥对为store的session
	router.POST("/register",M(), V.S2)
	router.POST("/login",M(), V.S1)

	router.POST("/test", M(),func(c *gin.Context) {
		username := c.PostForm("username")
		fmt.Println("you input right username", username)
		c.JSON(http.StatusOK, gin.H{
			"status": "S",
		})
	})

	router.Run(":9000")
}
