package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func response(c *gin.Context, status int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	})
}

func responseBadRequest(c *gin.Context, msg string) {
	response(c, 1, msg, nil)
}
func responseFatal(c *gin.Context, msg string) {
	response(c, 2, msg, nil)
}

func responseOK(c *gin.Context, data interface{}) {
	response(c, http.StatusOK, "", data)
}
