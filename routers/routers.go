package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由和main拆分开
func r0(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello matt!",
	})
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/r", r0)
	return r
}
