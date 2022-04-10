package main

import (
	"fmt"
	"study-gin/routers/r2"
	"study-gin/routers/r2/book"
)

func main() {
	// 加载多个APP的路由配置
	r2.Include(book.Routers)
	// 初始化路由
	r := r2.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
