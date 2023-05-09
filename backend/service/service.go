package service

import (
	"net/http"
	"server/service/dal"
	"server/service/dal/model"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Handler func(c *gin.Context) error

var db *gorm.DB

func StartUp(user, password, database_name string) error {
	tem, err := dal.NewDB(user, password, database_name)
	if err != nil {
		return err
	}
	db = tem
	return nil
}

func TestHandler(c *gin.Context) {
	offer := &model.Offer{
		RewardAmount:       100,
		CustomerId:         1,
		CategoryName:       "外卖",
		PickupLocationId:   1,
		DeliveryLocationId: 2,
		TimeLimit:          time.Now(),
	}
	offer_list := make([]model.Offer, 0)
	offer_list = append(offer_list, *offer)
	// offer_list = append(offer_list, *offer)
	c.JSON(http.StatusOK, offer_list)
}
