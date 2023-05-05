package service

import (
	"fmt"
	"server/service/dal/crud"
	"server/service/dal/model"
	"testing"
)

func TestRegister(t *testing.T) {
	StartUp("root", "123456", "pku_mutualhelper")
	db.Create(&model.User{
		NickName:     "Alice",
		FullName:     "Alice",
		PhoneNumber:  "1234567",
		PasswordHash: "password",
	})
	u := crud.FindUserByPhongNumber(db, "12345678")
	fmt.Println(u)
}
