package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("name")
		if cookie == "" {
			c.Abort()
			c.JSON(200, "err")
			return
		} else {
			c.Next()
		}

	}
}

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	//r.Use(MiddleWare())
	// 服务端要给客户端cookie
	r.GET("/l", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("name")
		if err != nil {
			c.SetCookie("name", "matt", 60, "/",
				"127.0.0.1", false, true)
		}
		fmt.Println(cookie, "l")
		c.JSON(http.StatusOK, "set")
	})

	r.GET("/h", MiddleWare(), func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, _ := c.Cookie("name")

		fmt.Printf("cookie的值是： %s\n", cookie)

		c.JSON(http.StatusOK, cookie)
	})
	r.Run(":8000")
}
