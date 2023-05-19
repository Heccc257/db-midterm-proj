package service

import (
	"net/http"
	"server/service/dal/crud"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AcceptOfferList(c *gin.Context) {
	if uid, err := strconv.Atoi(c.Param("uid")); err != nil {
		c.String(http.StatusBadRequest, "输入正确的uid")
	} else {
		accpet_offer_list := *crud.AcceptOfferListUser(uint(uid), db)
		c.JSON(http.StatusOK, accpet_offer_list)
	}

}
