package crud

import (
	"server/service/dal/model"

	"gorm.io/gorm"
)

func OfferListPending(db *gorm.DB) []model.Offer {
	offer_list := make([]model.Offer, 0)
	return offer_list
}
