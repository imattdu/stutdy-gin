package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func r11(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello user!",
	})
}

func LoadBook(e *gin.Engine) {
	e.GET("/r11", r11)
}
