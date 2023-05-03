package model

/*
create table Users
(
   user_id              int not null,
   nickname             varchar(20) not null,
   full_name            varchar(20) not null,
   phone_number         varchar(20) not null,
   password_hash        char(64) not null,
   order_count          int not null default 0,
   order_complaint_count int not null default 0,
   task_count           int not null default 0,
   task_complaint_count int not null default 0,
   primary key (user_id),
   key AK_u_phone (phone_number)
);

*/
type User struct {
	UserId              int    `gorm:"column:user_id;type:int(11);primary_key" json:"user_id"`
	Nickname            string `gorm:"column:nickname;type:varchar(20);NOT NULL" json:"nickname"`
	FullName            string `gorm:"column:full_name;type:varchar(20);NOT NULL" json:"full_name"`
	PhoneNumber         string `gorm:"column:phone_number;type:varchar(20);NOT NULL" json:"phone_number"`
	PasswordHash        string `gorm:"column:password_hash;type:char(64);NOT NULL" json:"password_hash"`
	OrderCount          int    `gorm:"column:order_count;type:int(11);default:0;NOT NULL" json:"order_count"`
	OrderComplaintCount int    `gorm:"column:order_complaint_count;type:int(11);default:0;NOT NULL" json:"order_complaint_count"`
	TaskCount           int    `gorm:"column:task_count;type:int(11);default:0;NOT NULL" json:"task_count"`
	TaskComplaintCount  int    `gorm:"column:task_complaint_count;type:int(11);default:0;NOT NULL" json:"task_complaint_count"`
}

func (m *User) TableName() string {
	return "user"
}
