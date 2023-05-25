package service

import (
	"fmt"
	"server/service/dal/crud"
	"server/service/dal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

/*
支持获取各种需求的offer_list
包括按照发布人索引，按照完成时间排序等等
*/

type OfferListShow struct {
	model.Offer
	CompleteTime time.Time `json:"complete_time"`
	WorkerID     uint      `json:"worker_id"`
}

func OfferListByUser(c *gin.Context) {
	fmt.Println("offer list uid", c.Param("uid"))
	if uid, err := strconv.Atoi(c.Param("uid")); err != nil {
		responseBadRequest(c, "请输入正确的uid")
		// c.String(http.StatusBadRequest, "输入正确的uid")
	} else {
		offer_list := *crud.OfferListUser(uint(uid), db)
		offer_list_show := make([]OfferListShow, len(offer_list))
		for i := 0; i < len(offer_list); i++ {
			offer_list_show[i].Offer = offer_list[i]
			if offer_list_show[i].Offer.OfferState != "pending" {
				crud.Offer2OfferShow(db, &offer_list_show[i].CompleteTime, &offer_list_show[i].WorkerID, offer_list_show[i].OfferId)
			}
		}
		// responseOK(c, offer_list)
		responseOK(c, offer_list_show)
		// c.JSON(http.StatusOK, offer_list)
	}
}

// 随机获取最多10条offer
func OfferListRandom(c *gin.Context) {
	var offer_list []model.Offer
	db.Where("offer_state = ?", "pending").Order("rand()").Limit(10).Find(&offer_list)
	responseOK(c, offer_list)
}

//按分类获取
func OfferListByCat(c *gin.Context) {
	fmt.Println("offer list by category", c.Param("category"))
	category := c.Param("category")
	if category == "" {
		responseBadRequest(c, "请输入正确的分类") //处理错误响应
	} else {
		var offerList []model.Offer
		db.Where("category_name = ? AND offer_state = ?", category, "pending").Find(&offerList)
		responseOK(c, offerList) //responseOK 替代原来的 c.JSON(http.StatusOK, offerList)
	}
}
