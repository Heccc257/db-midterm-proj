package model

/*

create table location
(
   location_id          int not null,
   building        VARCHAR(30) not null,
   floor VARCHAR(30),
   primary key (location_id)
);

*/
type Location struct {
	LocationId int    `gorm:"column:location_id;type:int(11);primary_key" json:"location_id"`
	Building   string `gorm:"column:building;type:varchar(30);NOT NULL" json:"building"`
	Floor      string `gorm:"column:floor;type:varchar(30)" json:"floor"`
}

func (m *Location) TableName() string {
	return "location"
}
