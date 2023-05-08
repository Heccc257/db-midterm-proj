package service

import (
	"net/http"
	"server/service/dal/crud"
	"server/service/dal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AcceptOffer(c *gin.Context) {
	accept_offer := model.AcceptOffer{}
	if offer_id, err := strconv.Atoi(c.Query("offer_id")); err != nil {
		c.String(http.StatusBadRequest, "请输入正确的offer_id")
	} else {
		accept_offer.OfferId = uint(offer_id)
	}
	if uid, err := strconv.Atoi(c.Query("uid")); err != nil {
		c.String(http.StatusBadRequest, "请输入正确的uid")
	} else {
		accept_offer.UserId = uint(uid)
	}
	accept_offer.CompleteTime = time.Now()

	// 事务同时插入accept_offer以及更新task_count
	tx := db.Begin()
	err1 := tx.Create(&accept_offer).Error
	err2 := crud.IncreseTaskCount(tx, accept_offer.UserId)
	err3 := crud.OfferState2InProgress(tx, accept_offer.OfferId)
	if err1 != nil || err2 != nil || err3 != nil {
		tx.Rollback()
		c.String(http.StatusInternalServerError, "接取订单失败")
		return
	}
	tx.Commit()
	c.String(http.StatusOK, "成功接到订单")
}
