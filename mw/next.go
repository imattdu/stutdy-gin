package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// MiddleWare1 定义中间 next方法会调用请求的方法
func MiddleWare1() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		// 设置变量到Context的key中，可以通过Get()取
		c.Set("request", "next")
		// 执行函数
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 注册中间件
	r.Use(MiddleWare1())
	// {}为了代码规范
	{
		r.GET("/next", func(c *gin.Context) {
			// 取值
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			//time.Sleep(2 * time.Second)
			// 页面接收
			c.JSON(200, gin.H{"request": req})
		})

	}
	r.Run(":8000")
}