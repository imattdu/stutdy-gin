package book

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func r21(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello r21!",
	})
}

func Routers(e *gin.Engine) {
	e.GET("/r21", r21)
}
