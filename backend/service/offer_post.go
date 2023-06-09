package service

import (
	"fmt"
	"server/service/dal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
	type Offer struct {
		OfferId            uint         `gorm:"column:offer_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"offer_id"`
		RewardAmount       float64      `gorm:"column:reward_amount;type:decimal(10,2);NOT NULL" json:"reward_amount"`
		CustomerId         uint         `gorm:"column:customer_id;type:int(11) unsigned;NOT NULL" json:"customer_id"`
		CategoryName       string       `gorm:"column:category_name;type:varchar(50);NOT NULL" json:"category_name"`
		PickupLocationId   uint         `gorm:"column:pickup_location_id;type:int(11) unsigned;NOT NULL" json:"pickup_location_id"`
		DeliveryLocationId uint         `gorm:"column:delivery_location_id;type:int(11) unsigned;NOT NULL" json:"delivery_location_id"`
		OfferState         string       `gorm:"column:offer_state;type:enum('pending','in_progress','completed');default:pending;NOT NULL" json:"offer_state"`
		CreatedAt          sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
		TimeLimit          time.Time    `gorm:"column:time_limit;type:datetime" json:"time_limit"`
	}
*/
type RequestOfferPost struct {
	CategoryName       string   `form:"category_name"`
	CustomerID         int64    `form:"customer_id"`
	DeliveryLocationID int64    `form:"delivery_location_id"`
	PickupLocationID   int64    `form:"pickup_location_id"`
	RewardAmount       *float64 `form:"reward_amount,omitempty"`
	TimeLimit          *string  `form:"time_limit,omitempty"`
}

func OfferPost(c *gin.Context) {
	var formOffer RequestOfferPost
	c.ShouldBind(&formOffer)
	offer := model.Offer{
		CustomerId:         uint(formOffer.CustomerID),
		CategoryName:       formOffer.CategoryName,
		PickupLocationId:   uint(formOffer.PickupLocationID),
		DeliveryLocationId: uint(formOffer.DeliveryLocationID),
	}

	token := c.Param("token")

	// 测试用的万能token
	if token != "access" {
		if uid, ok := tokens[token]; !ok || uid != offer.CustomerId {
			responseBadRequest(c, fmt.Sprintf("token验证失败 token: %s, uid: %d", token, offer.CustomerId))
			// c.String(http.StatusBadRequest, "token验证失败")
			return
		}
	}

	if formOffer.RewardAmount != nil {
		offer.RewardAmount = *formOffer.RewardAmount
	} else {
		offer.RewardAmount = -1
	}

	if formOffer.TimeLimit != nil {
		if time_limit, err := strconv.Atoi(*formOffer.TimeLimit); err != nil {
			responseBadRequest(c, "请输入正确的截止时间")
			// c.String(http.StatusBadRequest, "请输入正确的截止时间")
			return
		} else {
			offer.TimeLimit = time.Unix(time.Now().Unix()+int64(time_limit), 0)
		}
	} else {
		offer.TimeLimit = time.Unix(time.Now().Unix()-1, 0)
	}

	// 事务
	// 在发表offer的时候将对应的user的offer_count++
	tx := db.Begin()

	if err := tx.Model(&model.User{}).Where("user_id = ?", offer.CustomerId).Update("offer_count", gorm.Expr("offer_count + 1")).Error; err != nil {
		tx.Rollback()
		responseFatal(c, err.Error())
		// c.String(http.StatusInternalServerError, err.Error())
		return
	}

	if err := tx.Create(&offer).Error; err != nil {
		tx.Rollback()
		responseFatal(c, err.Error())
		return
	}

	tx.Commit()
	// c.JSON(http.StatusOK, gin.H{
	// 	"id": offer.OfferId,
	// })
	responseOK(c, gin.H{
		"id": offer.OfferId,
	})
}
