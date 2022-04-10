package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 多种响应方式
func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	// 2. 结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})

	// 重定向 url改变 不共享servlet
	r.GET("/d1", func(c *gin.Context) {

		c.Redirect(http.StatusMovedPermanently, "localhost:8000/d2")
	})

	// 转发
	r.GET("/d10", func(c *gin.Context) {

		c.Request.URL.Path = "/d2" //把请求的URL修改
		r.HandleContext(c)         //继续后续的处理
	})

	r.GET("/d2", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "d2", "status": 200})
	})

	r.Run(":8000")
}
