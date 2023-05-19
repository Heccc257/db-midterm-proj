package crud

import (
	"fmt"
	"server/service/dal/model"

	"gorm.io/gorm"
)

func FindAcceptOffer(db *gorm.DB, offer_id, uid uint, ao *model.AcceptOffer) error {
	db.Model(&model.AcceptOffer{}).Where("offer_id = ?", offer_id).Where("user_id = ?", uid).Find(ao)
	if ao.OfferId != offer_id || ao.UserId != uid {
		return fmt.Errorf("未找到对应表项")
	}
	return nil
}

func AcceptOfferListUser(uid uint, db *gorm.DB) *[]model.AcceptOffer {
	var accept_offer_list []model.AcceptOffer
	db.Model(&model.AcceptOffer{}).Where("user_id = ?", uid).Find(&accept_offer_list)
	return &accept_offer_list
}

func AcceptOfferModifyState(tx *gorm.DB, offer_id, uid uint, state string) error {
	var accept_offer model.AcceptOffer
	if tx.Model(&model.AcceptOffer{}).Where("offer_id = ?", offer_id).Where("uid = ?", uid).Find(&accept_offer); accept_offer.OfferId == 0 {
		return fmt.Errorf("未找到对应accept_offer")
	}
	accept_offer.TaskState = state
	tx.Save(&accept_offer)
	return nil
}
