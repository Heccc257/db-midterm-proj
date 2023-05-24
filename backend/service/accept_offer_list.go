package service

import (
	"server/service/dal/crud"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AcceptOfferList(c *gin.Context) {
	if uid, err := strconv.Atoi(c.Param("uid")); err != nil {
		responseBadRequest(c, "输入正确的uid")
		// c.String(http.StatusBadRequest, "输入正确的uid")
	} else {
		accpet_offer_list := *crud.AcceptOfferListUser(uint(uid), db)
		responseOK(c, accpet_offer_list)
		// c.JSON(http.StatusOK, accpet_offer_list)
	}

}
