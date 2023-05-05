package dal

import (
	"fmt"
	"log"
	"server/service/dal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(user, pass, dbname string) (db *gorm.DB, err error) {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// 先不打开具体的数据库
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local", user, pass)
	fmt.Println("dsn: ", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Exec(fmt.Sprintf("drop database if exists %s;", dbname))
	db.Exec(fmt.Sprintf("create database if not exists %s;", dbname))
	db.Exec(fmt.Sprintf("use %s;", dbname))
	if err := createTables(db); err != nil {
		return nil, err
	}
	return
}

func createTables(db *gorm.DB) error {
	log.Println("creating tables")
	err := db.AutoMigrate(
		&model.User{},
		&model.Complaint{},
		&model.Location{},
		&model.OfferCategory{},
		&model.Offer{},
		&model.AcceptOffer{},
		&model.UserRating{},
	)
	if err != nil {
		return err
	}
	model.AddCategories(db)
	log.Println("adding foreign key")
	db.Exec(`alter table offer ADD CONSTRAINT FK_OFFER_USER foreign key(customer_id) references user(user_id) on delete  CASCADE;`)
	return nil
}
