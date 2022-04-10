package main

import (
	"fmt"
	"study-gin/routers"
)

func main() {

	r := routers.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}

}
