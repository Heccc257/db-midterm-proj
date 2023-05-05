package model

type User struct {
	UserId              uint   `gorm:"column:user_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"user_id"`
	NickName            string `gorm:"column:nick_name;type:varchar(20);NOT NULL" json:"nick_name"`
	FullName            string `gorm:"column:full_name;type:varchar(20);NOT NULL" json:"full_name"`
	PhoneNumber         string `gorm:"column:phone_number;type:varchar(20);NOT NULL" json:"phone_number"`
	PasswordHash        string `gorm:"column:password_hash;type:char(64);NOT NULL" json:"password_hash"`
	OfferCount          int    `gorm:"column:offer_count;type:int(11);default:0;NOT NULL" json:"offer_count"`
	OfferComplaintCount int    `gorm:"column:offer_complaint_count;type:int(11);default:0;NOT NULL" json:"offer_complaint_count"`
	TaskCount           int    `gorm:"column:task_count;type:int(11);default:0;NOT NULL" json:"task_count"`
	TaskComplaintCount  int    `gorm:"column:task_complaint_count;type:int(11);default:0;NOT NULL" json:"task_complaint_count"`
}

func (m *User) TableName() string {
	return "user"
}
