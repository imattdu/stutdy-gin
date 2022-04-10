package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func t1(c *gin.Context) {

	c.String(200, fmt.Sprintf("hello t1"))
}

func t2(c *gin.Context) {
	c.String(200, fmt.Sprintf("hello t2"))
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	// m1.api 参数
	r.GET("/book/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		//截取/
		//action = strings.Trim(action, "/")

		// :一级目录 *多层目录
		//http://127.0.0.1:8000/book/n1/n2/n3/n4
		// name=n2 action=/n2/n3/n4
		c.String(http.StatusOK, name+" is "+action)
	})

	// m2.url 参数
	r.GET("/book1", func(c *gin.Context) {
		name := c.DefaultQuery("name", "matt")
		c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})

	// m3.表单参数
	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		username := c.PostForm("username")
		password := c.PostForm("password")
		// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
		c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})

	// m4上传文件
	r.POST("/upload0", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		// c.JSON(200, gin.H{"message": file.Header.Context})
		err = c.SaveUploadedFile(file, fmt.Sprintf("./res/%s", file.Filename))
		if err != nil {
			return
		}
		c.String(http.StatusOK, file.Filename)
	})

	// m4 对上传的文件进行限制
	r.POST("/upload1", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		//headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		//headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			c.String(http.StatusInternalServerError, "只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./res/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})

	// m4 上传多个文件 /upload/mul
	r.POST("/upload/mul", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, "./res/"+file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})

	// m5 路由组
	v1 := r.Group("/v1")
	// {} 是书写规范
	{
		v1.GET("/a1", t1)
		v1.GET("a2", t2)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/a1", t1)
		v2.POST("/a2", t2)
	}

	r.POST("/testJson", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
