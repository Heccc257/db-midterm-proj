package dal

import (
	"fmt"
	"server/service/dal/model"
	"testing"
	"time"
)

func TestNewDB(t *testing.T) {
	NewDB("root", "123456", "pku_mutualhelper")
	// db.Exec("drop database if exists bank;")
}

func TestCreateTables(t *testing.T) {
	// t := time.Now().Unix()
	// fmt.Println(time.Now(), time.Unix(time.Now().Unix()+1000000, 0))

	db, _ := NewDB("root", "123456", "pku_mutualhelper")
	db.Create(&model.User{
		NickName:     "Alice",
		FullName:     "Alice",
		PhoneNumber:  "1234567",
		PasswordHash: "password",
	})

	m := &model.Offer{
		RewardAmount:       100,
		CustomerId:         1,
		CategoryName:       "外卖",
		PickupLocationId:   1,
		DeliveryLocationId: 2,
		TimeLimit:          time.Now(),
	}
	db.Create(m)
	fmt.Println(m.CreatedAt.Time.Unix())
}
