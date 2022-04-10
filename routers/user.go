package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func r10(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello user!",
	})
}

func LoadUser(e *gin.Engine) {
	e.GET("/r10", r10)
}
