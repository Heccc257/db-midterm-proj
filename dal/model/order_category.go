package model

/*
create table order_category
(
   category_id    int unsigned AUTO_INCREMENT,
   category_name        varchar(50) not null,
   adviced_reward decimal(10,2) DEFAULT 100,
   adviced_timelimit_min int DEFAULT 100,
   primary key (category_id)
);
*/

type OrderCategory struct {
	CategoryId          uint    `gorm:"column:category_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"category_id"`
	CategoryName        string  `gorm:"column:category_name;type:varchar(50);NOT NULL" json:"category_name"`
	AdvicedReward       float64 `gorm:"column:adviced_reward;type:decimal(10,2);default:100" json:"adviced_reward"`
	AdvicedTimelimitMin int     `gorm:"column:adviced_timelimit_min;type:int(11);default:100" json:"adviced_timelimit_min"`
}

func (m *OrderCategory) TableName() string {
	return "order_category"
}
