package crud

import (
	"server/service/dal/model"

	"gorm.io/gorm"
)

type user_info struct {
	UserId              uint   `gorm:"column:user_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"user_id"`
	Nickname            string `gorm:"column:nickname;type:varchar(20);NOT NULL" json:"nickname"`
	FullName            string `gorm:"column:full_name;type:varchar(20);NOT NULL" json:"full_name"`
	PhoneNumber         string `gorm:"column:phone_number;type:varchar(20);NOT NULL" json:"phone_number"`
	OfferCount          int    `gorm:"column:offer_count;type:int(11);default:0;NOT NULL" json:"offer_count"`
	OfferComplaintCount int    `gorm:"column:offer_complaint_count;type:int(11);default:0;NOT NULL" json:"offer_complaint_count"`
	TaskCount           int    `gorm:"column:task_count;type:int(11);default:0;NOT NULL" json:"task_count"`
	TaskComplaintCount  int    `gorm:"column:task_complaint_count;type:int(11);default:0;NOT NULL" json:"task_complaint_count"`
}

func User2user_info(u *model.User) *user_info {
	return &user_info{
		UserId:              u.UserId,
		Nickname:            u.Nickname,
		FullName:            u.FullName,
		PhoneNumber:         u.PhoneNumber,
		OfferCount:          u.OfferCount,
		OfferComplaintCount: u.OfferComplaintCount,
		TaskCount:           u.TaskCount,
		TaskComplaintCount:  u.TaskComplaintCount,
	}
}

func FindUserByPhongNumber(db *gorm.DB, PhoneNumber string) *model.User {
	var u model.User
	db.Where("phone_number = ?", PhoneNumber).Find(&u)
	return &u
}
func FindUserByUID(db *gorm.DB, uid uint) *model.User {
	var u model.User
	db.Where("user_id = ?", uid).Find(&u)
	return &u
}

func InsertUser(db *gorm.DB, u *model.User) {
	db.Create(u)
}
