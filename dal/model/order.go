package model

import (
	"database/sql"
	"time"
)

/*
create table `order`
(
   order_id             int unsigned not null AUTO_INCREMENT,
   reward_amount        decimal(10,2) not null,
   customer_id          int unsigned not null,
   category          varchar(50) not null,
   pickup_location_id   int unsigned not null,
   delivery_location_id int unsigned not null,
	created_at timestamp NULL DEFAULT NULL,
   time_limit           Datetime,
   primary key (order_id)
);

*/

type Order struct {
	OrderId            uint         `gorm:"column:order_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"order_id"`
	RewardAmount       float64      `gorm:"column:reward_amount;type:decimal(10,2);NOT NULL" json:"reward_amount"`
	CustomerId         uint         `gorm:"column:customer_id;type:int(11) unsigned;NOT NULL" json:"customer_id"`
	Category           string       `gorm:"column:category;type:varchar(50);NOT NULL" json:"category"`
	PickupLocationId   uint         `gorm:"column:pickup_location_id;type:int(11) unsigned;NOT NULL" json:"pickup_location_id"`
	DeliveryLocationId uint         `gorm:"column:delivery_location_id;type:int(11) unsigned;NOT NULL" json:"delivery_location_id"`
	CreatedAt          sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	TimeLimit          time.Time    `gorm:"column:time_limit;type:datetime" json:"time_limit"`
}

func (m *Order) TableName() string {
	return "order"
}
