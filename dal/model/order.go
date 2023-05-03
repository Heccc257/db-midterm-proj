package model

import (
	"database/sql"
	"time"
)

/*
create table `order`
(
   order_id             int not null,
   reward_amount        decimal(10,2) not null,
   customer_id          int not null,
   category          varchar(50) not null,
   pickup_location_id   int not null,
   delivery_location_id int not null,
	created_at timestamp NULL DEFAULT NULL,
   time_limit           Datetime,
   primary key (order_id)
);

*/

type Order struct {
	OrderId            int          `gorm:"column:order_id;type:int(11);primary_key" json:"order_id"`
	RewardAmount       float64      `gorm:"column:reward_amount;type:decimal(10,2);NOT NULL" json:"reward_amount"`
	CustomerId         int          `gorm:"column:customer_id;type:int(11);NOT NULL" json:"customer_id"`
	Category           string       `gorm:"column:category;type:varchar(50);NOT NULL" json:"category"`
	PickupLocationId   int          `gorm:"column:pickup_location_id;type:int(11);NOT NULL" json:"pickup_location_id"`
	DeliveryLocationId int          `gorm:"column:delivery_location_id;type:int(11);NOT NULL" json:"delivery_location_id"`
	CreatedAt          sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	TimeLimit          time.Time    `gorm:"column:time_limit;type:datetime" json:"time_limit"`
}

func (m *Order) TableName() string {
	return "order"
}
