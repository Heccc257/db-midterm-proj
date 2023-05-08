package crud

import (
	"fmt"
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

func IncreseTaskCount(tx *gorm.DB, id uint) error {
	var user_info model.UserInfo
	if tx.Model(&model.UserInfo{}).Where("user_id = ?", id).Find(&user_info); user_info.UserId == 0 {
		return fmt.Errorf("未找到user")
	}
	user_info.TaskCount += 1
	tx.Save(&user_info)
	return nil
}
