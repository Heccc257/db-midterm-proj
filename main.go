package main

import (
	"flag"
	"fmt"
	"server/service"

	"github.com/gin-gonic/gin"
)

var (
	port          string
	addr          string
	user          string
	password      string
	database_name string
)

func main() {
	flag.StringVar(&addr, "addr", "127.0.0.1", "address to listen")
	flag.StringVar(&port, "port", "9999", "port to listen")
	flag.StringVar(&user, "u", "root", "mysql user")
	flag.StringVar(&password, "p", "123456", "mysql password")
	flag.StringVar(&database_name, "db_name", "pku_mutualhelper", "mysql database name")
	flag.Parse()

	// gin.SetMode(gin.ReleaseMode)
	S := gin.Default()
	service.StartUp(user, password, database_name)

	S.GET("/", service.TestHandler)

	S.POST("/register", service.Register)

	err := S.Run(addr + ":" + port)
	if err != nil {
		fmt.Println("服务器启动失败! ", err)
	}
}
