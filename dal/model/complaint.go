package model

import (
	"time"
)

/*
create table complaint
(
   complaint_id         int not null,
   complainant_id       int not null,
   defendant_id         int not null,
   order_id             int not null,
   complaint_reason     text not null,
   time                 timestamp not null,
   primary key (complaint_id)
);

*/

type Complaint struct {
	ComplaintId     int       `gorm:"column:complaint_id;type:int(11);primary_key" json:"complaint_id"`
	ComplainantId   int       `gorm:"column:complainant_id;type:int(11);NOT NULL" json:"complainant_id"`
	DefendantId     int       `gorm:"column:defendant_id;type:int(11);NOT NULL" json:"defendant_id"`
	OrderId         int       `gorm:"column:order_id;type:int(11);NOT NULL" json:"order_id"`
	ComplaintReason string    `gorm:"column:complaint_reason;type:text;NOT NULL" json:"complaint_reason"`
	Time            time.Time `gorm:"column:time;type:timestamp;NOT NULL" json:"time"`
}

func (m *Complaint) TableName() string {
	return "complaint"
}
