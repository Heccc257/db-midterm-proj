package main

import (
	"flag"
	"fmt"
	"net/http"

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
	S.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "OK"})
	})

	S.GET("/:name", func(c *gin.Context) {
		c.String(http.StatusOK, "your name is %s\n", c.Param("name"))
		fmt.Println("query", c.Query("token"))
	})
	S.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test test")
	})

	err := S.Run(addr + ":" + port)
	if err != nil {
		fmt.Println("服务器启动失败! ", err)
	}
}
