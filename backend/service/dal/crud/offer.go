package crud

import (
	"fmt"
	"server/service/dal/model"
	"time"

	"gorm.io/gorm"
)

func OfferListUser(uid uint, db *gorm.DB) *[]model.Offer {
	var offer_list []model.Offer
	db.Model(&model.Offer{}).Where("customer_id = ?", uid).Find(&offer_list)
	return &offer_list
}
func OfferModifyState(tx *gorm.DB, offer_id uint, state string) error {
	var offer model.Offer
	if tx.Model(&model.Offer{}).Where("offer_id = ?", offer_id).Find(&offer); offer.OfferId == 0 {
		return fmt.Errorf("未找到对应offer")
	}
	offer.OfferState = state
	tx.Save(&offer)
	return nil
}

func Offer2OfferShow(db *gorm.DB, complete_time *time.Time, worker_id *uint, offer_id uint) {
	var accept_offer model.AcceptOffer
	db.Model(&model.AcceptOffer{}).Where("offer_id = ?", offer_id).Find(&accept_offer)
	*complete_time = accept_offer.CompleteTime
	*worker_id = accept_offer.UserId
}
