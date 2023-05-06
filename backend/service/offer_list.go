package service

import (
	"net/http"
	"server/service/dal/crud"
	"server/service/dal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
支持获取各种需求的offer_list
包括按照发布人索引，按照完成时间排序等等
*/

func OfferListByUser(c *gin.Context) {
	if uid, err := strconv.Atoi(c.Query("uid")); err != nil {
		c.String(http.StatusBadRequest, "输入正确的uid")
	} else {
		offer_list := *crud.OfferListUser(uint(uid), db)
		c.JSON(http.StatusOK, offer_list)
	}
}

// 随机获取最多10条offer
func OfferListRandom(c *gin.Context) {
	var offerList []model.Offer
	db.Where("offer_state = ?", "pending").Order("rand()").Limit(10).Find(&offerList)
	c.JSON(http.StatusOK, offerList)
}
