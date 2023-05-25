package service

import (
	"fmt"
	"server/service/dal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AcceptOfferList(c *gin.Context) {
	if uid, err := strconv.Atoi(c.Param("uid")); err != nil {
		responseBadRequest(c, "输入正确的uid")
		// c.String(http.StatusBadRequest, "输入正确的uid")
	} else {
		var offer_list []model.Offer

		// offer_id 相等
		accpet_offer_id := db.Table("accept_offer").
			Select("offer_id").
			Where("user_id = ?", uid)
		db.Table("(?) as t1", accpet_offer_id).
			Joins("NATURAL JOIN offer as t2").
			Select("*").Scan(&offer_list)
		// db.Table("offer").
		// 	Joins("NATURAL JOIN accept_offer").
		// 	Where("user_id = ?", uid).
		// 	Select("*").Scan(&offer_list)
		fmt.Println(offer_list)

		// accpet_offer_list := *crud.AcceptOfferListUser(uint(uid), db)
		// responseOK(c, accpet_offer_list)
		responseOK(c, offer_list)
		// c.JSON(http.StatusOK, accpet_offer_list)
	}

}

/*
select * from
accept_offer
NATURAL JOIN offer
where user_id = 1;
*/
