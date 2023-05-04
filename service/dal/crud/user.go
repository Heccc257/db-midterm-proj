package crud

import (
	"server/service/dal/model"

	"gorm.io/gorm"
)

func FindUserByPhongNumber(db gorm.DB, PhoneNumber string) *model.User {
	var u model.User
	db.Where("phone_number = ?", PhoneNumber).Find(&u)
	return &u
}
