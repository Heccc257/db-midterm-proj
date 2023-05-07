package model

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Offer struct {
	OfferId uint `gorm:"column:offer_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"offer_id"`
	// RewardAmount       float64      `gorm:"column:reward_amount;type:decimal(10,2);NOT NULL" json:"reward_amount"`
	RewardAmount       float64      `gorm:"column:reward_amount;type:decimal(10,2);NOT NULL" json:"reward_amount" form:"reward_amount"`
	CustomerId         uint         `gorm:"column:customer_id;type:int(11) unsigned;NOT NULL" json:"customer_id"`
	CategoryName       string       `gorm:"column:category_name;type:varchar(50);NOT NULL" json:"category_name"`
	PickupLocationId   uint         `gorm:"column:pickup_location_id;type:int(11) unsigned;NOT NULL" json:"pickup_location_id"`
	DeliveryLocationId uint         `gorm:"column:delivery_location_id;type:int(11) unsigned;NOT NULL" json:"delivery_location_id"`
	OfferState         string       `gorm:"column:offer_state;type:enum('pending','in_progress','completed');default:pending;NOT NULL" json:"offer_state"`
	CreatedAt          sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	TimeLimit          time.Time    `gorm:"column:time_limit;type:datetime" json:"time_limit"`
}

func (m *Offer) TableName() string {
	return "offer"
}

func RegisterOfferTrigger(db *gorm.DB) {
	trigger := "CREATE TRIGGER set_default_offer BEFORE INSERT ON `offer`\n" +
		`
	FOR EACH ROW
	BEGIN

		DECLARE reward FLOAT;
		DECLARE time_limit_min int;
		SELECT adviced_reward, adviced_timelimit_min INTO reward, time_limit_min FROM offer_category WHERE category_name = NEW.category_name;

		IF NEW.reward_amount < 1 THEN
			SET NEW.reward_amount = reward;
		END IF;
		IF NEW.time_limit < NOW() THEN
			SET NEW.time_limit = DATE_ADD(NOW(), INTERVAL time_limit_min MINUTE);
  		END IF;
	END;
	`
	fmt.Println("add trigger")
	db.Exec(trigger)
	// db.Exec("ALTER TRIGGER set_default_offer ENABLE;")
}

// func (m *Offer) AddForeignKey(db *gorm.DB) error {
// 	if err := db.Exec(`
// 	alter table offer ADD CONSTRAINT FK_OFFER_USER
// 	foreign key(offer_id) references user(user_id) on delete  CASCADE;
// 	`).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

/*
alter table offer ADD CONSTRAINT FK_OFFER_USER foreign key(offer_id) references user(user_id) on delete  CASCADE;

*/
