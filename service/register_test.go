package service

import (
	"fmt"
	"server/service/dal/crud"
	"server/service/dal/model"
	"testing"
)

func TestRegister(t *testing.T) {
	StartUp()
	db.Create(&model.User{
		Nickname:     "Alice",
		FullName:     "Alice",
		PhoneNumber:  "1234567",
		PasswordHash: "password",
	})
	u := crud.FindUserByPhongNumber(db, "12345678")
	fmt.Println(u)
}
