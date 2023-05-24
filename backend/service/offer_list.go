package service

import (
	"fmt"
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
	fmt.Println("offer list uid", c.Param("uid"))
	if uid, err := strconv.Atoi(c.Param("uid")); err != nil {
		responseBadRequest(c, "请输入正确的uid")
		// c.String(http.StatusBadRequest, "输入正确的uid")
	} else {
		offer_list := *crud.OfferListUser(uint(uid), db)
		responseOK(c, offer_list)
		// c.JSON(http.StatusOK, offer_list)
	}
}

// 随机获取最多10条offer
func OfferListRandom(c *gin.Context) {
	var offer_list []model.Offer
	db.Where("offer_state = ?", "pending").Order("rand()").Limit(10).Find(&offer_list)
	responseOK(c, offer_list)
}
