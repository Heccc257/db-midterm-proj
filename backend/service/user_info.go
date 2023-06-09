package service

import (
	"fmt"
	"server/service/dal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func User_info(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		responseBadRequest(c, "请输入正确格式的uid")
		// c.String(http.StatusBadRequest, "请输入正确格式的uid")
		return
	}
	fmt.Println("uid = ", uid)
	var user_info model.UserInfo
	db.Raw("CALL GetUserByID(?);", uid).Scan(&user_info)

	if user_info.UserId == uint(uid) {
		responseOK(c, user_info)
		// c.JSON(http.StatusOK, user_info)
		return
	}
	responseBadRequest(c, "未找到用户,请检查uid是否正确")
	// c.String(http.StatusNotFound, "未找到用户,请检查uid是否正确")
}
