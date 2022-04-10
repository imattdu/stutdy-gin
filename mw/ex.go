package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func spendTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		t2 := time.Since(t)
		fmt.Println("花费时间：", t2)
	}
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(spendTime())
	// {}为了代码规范
	{
		r.GET("/e1", func(c *gin.Context) {
			c.JSON(200, gin.H{"request": "matt"})
		})

		r.GET("/e2", func(c *gin.Context) {
			c.JSON(200, gin.H{"request": "17"})
		})

	}
	r.Run(":8000")
}
