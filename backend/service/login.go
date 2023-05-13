package service

import (
	"fmt"
	"net/http"
	"server/service/dal/model"

	"github.com/gin-gonic/gin"
)

var tokens map[string]uint = make(map[string]uint, 0)

func Login(c *gin.Context) {
	phone_number := c.PostForm("phone_number")
	password := c.PostForm("password")
	fmt.Println(phone_number, password)
	var user model.User
	db.Model(&model.User{}).Where("phone_number = ?", phone_number).Find(&user)
	if user.UserId == 0 {
		c.String(http.StatusBadRequest, "该手机号还未注册")
		return
	}
	if user.PasswordHash != password {
		c.String(http.StatusBadRequest, "密码错误")
		return
	}
	token := password + "*" + phone_number
	tokens[token] = user.UserId
	c.JSON(http.StatusOK, gin.H{
		"uid":   user.UserId,
		"token": token,
	})
}
