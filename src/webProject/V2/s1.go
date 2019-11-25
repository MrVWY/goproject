package V2

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Getv2(c *gin.Context)  {
	c.String(http.StatusOK,"V2 Hello")
}