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
	dsn := fmt.Sprintf("/?charset=utf8mb4&parseTime=True&loc=Local")
	fmt.Println("dsn: ", dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
		return
	} else {
		fmt.Println("连接数据库成功")
	}
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

	// 插入原始数据
	model.AddCategories(db)
	model.AddLocations(db)

	// 添加外键
	log.Println("adding foreign key")
	// db.Exec(`alter table offer ADD CONSTRAINT FK_OFFER_USER foreign key(customer_id) references user(user_id) on delete  CASCADE;`)
	db.Exec(`
	alter table accept_offer add constraint FK_Reference_15 foreign key (user_id)
      references user (user_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table accept_offer add constraint FK_Reference_16 foreign key (offer_id)
      references offer (offer_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table complaint add constraint FK_Reference_6 foreign key (complainant_id)
      references user (user_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table complaint add constraint FK_Reference_7 foreign key (defendant_id)
      references user (user_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table complaint add constraint FK_Reference_8 foreign key (offer_id)
      references offer (offer_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table offer add constraint FK_Reference_1 foreign key (customer_id)
      references user (user_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table offer add constraint FK_Reference_19 foreign key (category_name)
      references offer_category (category_name) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table offer add constraint FK_Reference_4 foreign key (pickup_location_id)
      references location (location_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table offer add constraint FK_Reference_5 foreign key (delivery_location_id)
      references location (location_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table user_rating add constraint FK_Reference_10 foreign key (rater_id)
      references user (user_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table user_rating add constraint FK_Reference_11 foreign key (offer_id)
      references offer (offer_id) on delete restrict on update restrict;
	`)
	db.Exec(`
	alter table user_rating add constraint FK_Reference_9 foreign key (rated_user_id)
      references user (user_id) on delete restrict on update restrict;
	`)

	// 添加视图
	model.CreateUserView(db)

	// 添加trigger
	model.RegisterOfferTrigger(db)

	// 添加存储过程
	db.Exec(`
		CREATE PROCEDURE GetUserByID(IN uid INT)
		BEGIN
			SELECT *
			FROM user_info
			WHERE user_id = uid;
		END
	`)

	// 修改时区
	log.Println("修改时区为UTC+8:00")
	db.Exec("set time_zone = '+8:00';")
	return nil
}
