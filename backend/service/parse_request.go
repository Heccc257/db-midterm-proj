package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Form2Map(c *gin.Context) *map[string]string {
	m := make(map[string]string)
	// 需要一次读取操作将form加载到内存中
	c.PostForm("load post form into memory")
	form := c.Request.PostForm
	for key, val := range form {
		m[key] = val[0]
		fmt.Println("val ", val[0])
	}
	return &m
}
