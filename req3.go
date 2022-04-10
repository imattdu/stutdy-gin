package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Name string `form:"name" json:"name" uri:"name" binding:"required"`
	Id   int    `form:"id" json:"id" uri:"id" binding:"required"`
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// JSON绑定
	r.GET("/req3", func(c *gin.Context) {
		// 声明接收的变量
		var b Book
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.ShouldBindQuery(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id": b.Id, "name": b.Name})
	})
	r.Run(":8000")

}
