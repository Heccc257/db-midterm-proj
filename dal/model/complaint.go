package model

import (
	"time"
)

/*
create table complaint
(
   complaint_id         int unsigned AUTO_INCREMENT not null,
   complainant_id       int unsigned not null,
   defendant_id         int unsigned not null,
   order_id             int unsigned not null,
   complaint_reason     text not null,
   time                 timestamp not null,
   primary key (complaint_id)
);
*/

type Complaint struct {
	ComplaintId     uint      `gorm:"column:complaint_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"complaint_id"`
	ComplainantId   uint      `gorm:"column:complainant_id;type:int(11) unsigned;NOT NULL" json:"complainant_id"`
	DefendantId     uint      `gorm:"column:defendant_id;type:int(11) unsigned;NOT NULL" json:"defendant_id"`
	OrderId         uint      `gorm:"column:order_id;type:int(11) unsigned;NOT NULL" json:"order_id"`
	ComplaintReason string    `gorm:"column:complaint_reason;type:text;NOT NULL" json:"complaint_reason"`
	Time            time.Time `gorm:"column:time;type:timestamp;NOT NULL" json:"time"`
}

func (m *Complaint) TableName() string {
	return "complaint"
}
