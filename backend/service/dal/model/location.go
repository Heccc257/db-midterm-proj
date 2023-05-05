package model

type Location struct {
	LocationId uint   `gorm:"column:location_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"location_id"`
	Building   string `gorm:"column:building;type:varchar(30);NOT NULL" json:"building"`
	Floor      string `gorm:"column:floor;type:varchar(30)" json:"floor"`
}

func (m *Location) TableName() string {
	return "location"
}
