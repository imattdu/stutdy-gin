package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"study-gin/routers"
)

func main() {

	r := gin.Default()
	routers.LoadUser(r)
	routers.LoadBook(r)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
