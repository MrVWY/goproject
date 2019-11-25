package view

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	Z "microservives/emamples/mianrouter/model"
	b "microservives/emamples/mianrouter/model/db"
)

const times  = 100

func S1(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("mysession") == nil{
		c.JSON(http.StatusOK,gin.H{
			"status":"The user is already login , redict now",
		})
		c.Abort()// 终止
		return
	}
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	db := b.Initialization()
	defer db.Close()
	err, user := b.R(db,&Z.User{
		Username: username,
		Password: pwd,
	})
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"The user is not exitis",
		})
		c.Abort()// 终止
		return
	}
	session.Set(username,pwd)
	session.Save()
	c.JSON(http.StatusOK,gin.H{
		"user":user,
	})
}


func S2(c *gin.Context)  {
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	db := b.Initialization()
	defer db.Close()
	err ,z := b.R(db,&Z.User{
		Username: username,
		Password: pwd,
	})
	fmt.Println(z,err)
	if err == nil{
		if z.Username == username{
			c.JSON(http.StatusBadRequest,gin.H{
				"status": "fail2",
			})
			c.Abort()// 终止
			return
		}else {
			b.C(db,&Z.User{
				Username: username,
				Password: pwd,
			})
			c.JSON(http.StatusOK,gin.H{
				"status": "successfully",
			})
			c.Abort()// 终止
			return
		}
	}else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status": "fail",
		})
	}
}

func S3(c *gin.Context)  {
	username := c.PostForm("username")
	pwd := c.PostForm("password")
	db := b.Initialization()
	defer db.Close()
	//根据username查询
	err, user := b.R(db,&Z.User{
		Username: username,
	})
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"The user is not exitis",
		})
	}
	if pwd == user.Password {
		//z := &N.User{
		//	Username: user.Username,
		//	Password: user.Password,
		//}
		//d := time.Duration(times)*time.Second
		//token, err := N.JwtGenerateToken(z,d)
		if err != nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"status":"The create token is fail",
			})
			c.Abort()// 终止
			return
		}else {
			c.JSON(http.StatusOK,gin.H{
				"user":username,
				//"token": token,
			})
			c.Abort()// 终止
			return
		}
	}else {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"The passwork is not match",
		})
		c.Abort()// 终止
		return
	}
}