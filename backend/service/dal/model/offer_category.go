package model

import "gorm.io/gorm"

type OfferCategory struct {
	CategoryId          uint    `gorm:"column:category_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"category_id"`
	CategoryName        string  `gorm:"column:category_name;type:varchar(50);NOT NULL;uniqueIndex" json:"category_name"`
	AdvicedReward       float64 `gorm:"column:adviced_reward;type:decimal(10,2);default:100" json:"adviced_reward"`
	AdvicedTimelimitMin int     `gorm:"column:adviced_timelimit_min;type:int(11);default:100" json:"adviced_timelimit_min"`
}

func (m *OfferCategory) TableName() string {
	return "offer_category"
}

func AddCategories(db *gorm.DB) {
	categories := []OfferCategory{
		{CategoryName: "拿外卖", AdvicedReward: 66666.0, AdvicedTimelimitMin: 60},
		{CategoryName: "跑腿", AdvicedReward: 30000.0, AdvicedTimelimitMin: 120},
	}
	db.Create(categories)
}

/*
{category_id: 1, category_name: "拿外卖", adviced_reward: 66666.0, adviced_timelimit_min: 60},
{category_id: 2, category_name: "跑腿", adviced_reward: 30000.0, adviced_timelimit_min: 120},


*/
