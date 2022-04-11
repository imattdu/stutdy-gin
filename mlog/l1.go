package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁用控制台日志颜色,日志写到文件的时候,不需要打开控制台日志颜色
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	r.GET("/l", func(c *gin.Context) {
		c.String(200, "l")
	})
	r.Run()
}
