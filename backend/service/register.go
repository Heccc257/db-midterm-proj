package service

import (
	"fmt"
	"log"
	"net/http"
	"server/service/dal/crud"
	"server/service/dal/model"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	fmt.Println("register")
	if c.PostForm("phone_number") == "" {
		c.String(http.StatusOK, "phone_number can't be empty")
		return
	}
	if c.PostForm("full_name") == "" {
		c.String(http.StatusOK, "full_name can't be empty")
		return
	}
	if c.PostForm("nickname") == "" {
		c.String(http.StatusOK, "nickname can't be empty")
		return
	}
	if c.PostForm("password_hash") == "" {
		c.String(http.StatusOK, "password_hash can't be empty")
		return
	}
	phone_number := c.PostForm("phone_number")
	log.Println("phone number = ", phone_number)
	u := crud.FindUserByPhongNumber(db, phone_number)
	if u != nil && u.UserId != 0 {
		fmt.Println("user: ", u)
		c.String(http.StatusOK, "user exist!")
		return
	}
	fmt.Println("full name = ", c.PostForm("full"))
	user := model.User{
		NickName:     c.PostForm("nick_name"),
		FullName:     c.PostForm("full_name"),
		PhoneNumber:  c.PostForm("phone_number"),
		PasswordHash: c.PostForm("password_hash"),
	}
	crud.InsertUser(db, &user)
	c.String(http.StatusOK, "register succeed and your uid is %d", user.UserId)
}
