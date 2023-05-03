package model

/*

create table location
(
   location_id          int unsigned AUTO_INCREMENT not null,
   building        VARCHAR(30) not null,
   floor 			VARCHAR(30),
   primary key (location_id)
);

*/
type Location struct {
	LocationId uint   `gorm:"column:location_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"location_id"`
	Building   string `gorm:"column:building;type:varchar(30);NOT NULL" json:"building"`
	Floor      string `gorm:"column:floor;type:varchar(30)" json:"floor"`
}

func (m *Location) TableName() string {
	return "location"
}
