package service

import (
	"net/http"
	"server/service/dal/model"

	"github.com/gin-gonic/gin"
)

/*
支持获取各种需求的offer_list
包括按照发布人索引，按照完成时间排序等等
*/

func offerListByUser(c *gin.Context) {

}

// 随机获取最多10条offer
func OfferListRandom(c *gin.Context) {
	var offerList []model.Offer
	db.Order("rand()").Limit(10).Find(&offerList)
	c.JSON(http.StatusOK, offerList)
}
