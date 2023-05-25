package service

import (
	"fmt"
	"log"
	"server/service/dal/model"

	"github.com/gin-gonic/gin"
)

var tokens map[string]uint = make(map[string]uint, 0)

func Login(c *gin.Context) {
	phone_number := c.PostForm("phone_number")
	password := c.PostForm("password")
	fmt.Println("login: ", phone_number, password)
	var user model.User
	fmt.Println("login phone_number = ", phone_number)
	db.Model(&model.User{}).Where("phone_number = ?", phone_number).Find(&user)
	if user.UserId == 0 {
		responseBadRequest(c, "该手机号还未注册")
		// c.String(http.StatusBadRequest, "该手机号还未注册")
		return
	}
	if user.PasswordHash != password {
		responseBadRequest(c, "密码错误")
		// c.String(http.StatusBadRequest, "密码错误")
		return
	}
	token := password + "*" + phone_number
	tokens[token] = user.UserId
	// c.JSON(http.StatusOK, gin.H{
	// 	"uid":   user.UserId,
	// 	"token": token,
	// })
	log.Printf("login: uid: %d ; token:%s \n", user.UserId, token)
	responseOK(c, gin.H{
		"uid":   user.UserId,
		"token": token,
	})
}
