package service

import (
	"net/http"
	"server/service/dal/crud"
	"strconv"

	"github.com/gin-gonic/gin"
)

func User_info(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.String(http.StatusBadRequest, "请输入正确格式的uid")
		return
	}
	u := crud.FindUserByUID(db, uint(uid))
	if u.UserId == uint(uid) {
		user_info := crud.User2user_info(u)
		c.JSON(http.StatusOK, user_info)
		return
	}
	c.String(http.StatusNotFound, "未找到用户,请检查uid是否正确")
}
