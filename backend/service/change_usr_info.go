package service

import (
	"net/http"
	"server/service/dal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ChangeUserInfo(c *gin.Context) {
	// Get the user id from the parameters
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid user id")
		return
	}

	// Find the user with the provided user id
	var user model.User
	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		c.String(http.StatusBadRequest, "User not found")
		return
	}

	// Extract the new user information from the request
	var newUserInfo model.UserInfo
	if err := c.ShouldBindJSON(&newUserInfo); err != nil {
		c.String(http.StatusBadRequest, "Invalid input")
		return
	}

	// Change the user's information if new information is provided
	if newUserInfo.NickName != "" {
		user.NickName = newUserInfo.NickName
	}
	if newUserInfo.FullName != "" {
		user.FullName = newUserInfo.FullName
	}
	if newUserInfo.PhoneNumber != "" {
		user.PhoneNumber = newUserInfo.PhoneNumber
	}

	// Save the changes to the database
	if err := db.Save(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, "Failed to update user info")
		return
	}

	c.String(http.StatusOK, "User info updated successfully")
}
