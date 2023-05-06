package crud

import (
	"server/service/dal/model"

	"gorm.io/gorm"
)

func OfferListUser(uid uint, db *gorm.DB) *[]model.Offer {
	var offer_list []model.Offer
	db.Model(&model.Offer{}).Where("customer_id = ?", uid).Find(&offer_list)
	return &offer_list
}
