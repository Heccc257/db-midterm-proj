package service

import (
	"fmt"
	"server/service/dal/crud"
	"server/service/dal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CompleteOffer(c *gin.Context) {

	query := model.AcceptOffer{}
	if offer_id, err := strconv.Atoi(c.Query("offer_id")); err != nil {
		responseBadRequest(c, "请输入正确的offer_id")
		// c.String(http.StatusBadRequest, "请输入正确的offer_id")
		return
	} else {
		query.OfferId = uint(offer_id)
	}
	if uid, err := strconv.Atoi(c.Query("uid")); err != nil {
		responseBadRequest(c, "请输入正确的uid")
		// c.String(http.StatusBadRequest, "请输入正确的uid")
		return
	} else {
		query.UserId = uint(uid)
	}

	token := c.Param("token")
	if token != "access" {
		fmt.Println("token = ", token)
		if uid, ok := tokens[token]; !ok || uid != query.UserId {
			responseBadRequest(c, "token验证失败")
			// c.String(http.StatusBadRequest, "token验证失败")
			return
		}
	}

	tx := db.Begin()

	var accept_offer model.AcceptOffer
	// 必须确保找到了对应的accept_offer
	err1 := crud.FindAcceptOffer(tx, query.OfferId, query.UserId, &accept_offer)

	err2 := crud.OfferModifyState(tx, accept_offer.OfferId, "completed")

	accept_offer.CompleteTime = time.Now()
	accept_offer.TaskState = "completed"
	err3 := tx.Save(&accept_offer).Error

	if err1 != nil || err2 != nil || err3 != nil {
		tx.Rollback()
		responseFatal(c, "完成订单失败")
		// c.String(http.StatusInternalServerError, "完成订单失败")
		return
	}
	tx.Commit()
	responseOK(c, "成功完成订单")
	// c.String(http.StatusOK, "成功完成订单")
}
