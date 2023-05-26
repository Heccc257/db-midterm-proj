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
	// fmt.Println("phone_number: ", c.PostForm("phone_number"))
	// fmt.Println("nick_name: ", c.PostForm("nick_name"))
	if c.PostForm("phone_number") == "" {
		responseBadRequest(c, "phone_number can't be empty")
		// c.String(http.StatusBadRequest, "phone_number can't be empty")
		return
	}
	if c.PostForm("full_name") == "" {
		responseBadRequest(c, "full_name can't be empty")
		// c.String(http.StatusBadRequest, "full_name can't be empty")
		return
	}
	if c.PostForm("nick_name") == "" {
		responseBadRequest(c, "nickname can't be empty")
		// c.String(http.StatusBadRequest, "nickname can't be empty")
		return
	}
	if c.PostForm("password_hash") == "" {
		responseBadRequest(c, "password_hash can't be empty")
		// c.String(http.StatusBadRequest, "password_hash can't be empty")
		return
	}
	phone_number := c.PostForm("phone_number")

	// if len(phone_number) != 11 {
	// 	//c.String(http.StatusBadRequest, "Phone number should be 11 digits")
	// 	responseBadRequest(c, "Phone number should be 11 digits")
	// 	return
	// }
	log.Println("phone number = ", phone_number)

	nickName := c.PostForm("nick_name")
	if len(nickName) > 25 {
		//c.String(http.StatusBadRequest, "NickName should be within 25 characters")
		responseBadRequest(c, "NickName should be within 25 characters")
		return
	}

	fullName := c.PostForm("full_name")
	if len(fullName) > 25 {
		//c.String(http.StatusBadRequest, "FullName should be within 25 characters")
		responseBadRequest(c, "FullName should be within 25 characters")
		return
	}

	u := crud.FindUserInfoByPhongNumber(db, phone_number)
	if u != nil && u.UserId != 0 {
		fmt.Println("user: ", u)
		responseBadRequest(c, "phone number has been used!")
		// c.String(http.StatusOK, "phone number has been used!")
		return
	}
	user := model.User{
		NickName:     c.PostForm("nick_name"),
		FullName:     c.PostForm("full_name"),
		PhoneNumber:  c.PostForm("phone_number"),
		PasswordHash: c.PostForm("password_hash"),
	}
	crud.InsertUser(db, &user)

	response(c, http.StatusOK, fmt.Sprintf("register succeed and your uid is %d", user.UserId), nil)
	// c.String(http.StatusOK, "register succeed and your uid is %d", user.UserId)
}
