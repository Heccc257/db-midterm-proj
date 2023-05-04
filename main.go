package main

import (
	"flag"
	"fmt"
	"server/service"

	"github.com/gin-gonic/gin"
)

var (
	port string
	addr string
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1", "address to listen")
	flag.StringVar(&port, "p", "9999", "port to listen")
	flag.Parse()

	// gin.SetMode(gin.ReleaseMode)
	S := gin.Default()
	service.StartUp()

	S.GET("/", service.TestHandler)

	S.POST("/register", service.Register)

	err := S.Run(addr + ":" + port)
	if err != nil {
		fmt.Println("服务器启动失败! ", err)
	}
}
