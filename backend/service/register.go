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
		c.String(http.StatusBadRequest, "phone_number can't be empty")
		return
	}
	if c.PostForm("full_name") == "" {
		c.String(http.StatusBadRequest, "full_name can't be empty")
		return
	}
	if c.PostForm("nick_name") == "" {
		c.String(http.StatusBadRequest, "nickname can't be empty")
		return
	}
	if c.PostForm("password_hash") == "" {
		c.String(http.StatusBadRequest, "password_hash can't be empty")
		return
	}
	phone_number := c.PostForm("phone_number")
	if len(phoneNumber) != 11 {
		c.String(http.StatusBadRequest, "Phone number should be 11 digits")
		return
	}
	log.Println("phone number = ", phone_number)

	nickName := c.PostForm("nick_name")
	if len(nickName) > 25 {
		c.String(http.StatusBadRequest, "NickName should be within 25 characters")
		return
	}

	fullName := c.PostForm("full_name")
	if len(fullName) > 25 {
		c.String(http.StatusBadRequest, "FullName should be within 25 characters")
		return
	}

	u := crud.FindUserInfoByPhongNumber(db, phone_number)
	if u != nil && u.UserId != 0 {
		fmt.Println("user: ", u)
		c.String(http.StatusOK, "phone number has been used!")
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
