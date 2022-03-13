package main

import (
	"context"
	"fmt"
	"github.com/guangtangchen/common/logs"
	"github.com/guangtangchen/local_test/model"
	"time"
)

func main() {
	//fmt.Println("hello world")
	//
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//for i := 0; i < 100; i++ {
	//	fmt.Println("i = ", i)
	//	time.Sleep(1 * time.Second)
	//}
	ctx := context.Background()
	logs.Init("test")
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("i = ", i)
			time.Sleep(1 * time.Second)
		}
	}()
	model.Init(ctx)

	//resp, err := model.Set(ctx, "test1", time.Now().String(), -1)
	//fmt.Println(resp, err)
	logs.CtxInfo(ctx, "hahah")
	time.Sleep(10 * time.Minute)
}
