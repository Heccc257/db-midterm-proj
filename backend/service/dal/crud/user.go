package crud

import (
	"server/service/dal/model"

	"gorm.io/gorm"
)

func FindUserInfoByPhongNumber(db *gorm.DB, PhoneNumber string) *model.UserInfo {
	var u model.UserInfo
	db.Where("phone_number = ?", PhoneNumber).Find(&u)
	return &u
}

func FindUserInfoByUID(db *gorm.DB, uid uint) *model.UserInfo {
	var u model.UserInfo
	db.Where("user_id = ?", uid).Find(&u)
	return &u
}

func InsertUser(db *gorm.DB, u *model.User) {
	db.Create(u)
}
