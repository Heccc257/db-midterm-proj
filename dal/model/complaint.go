package model

import (
	"time"
)

type Complaint struct {
	ComplaintId     uint      `gorm:"column:complaint_id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"complaint_id"`
	ComplainantId   uint      `gorm:"column:complainant_id;type:int(11) unsigned;NOT NULL" json:"complainant_id"`
	DefendantId     uint      `gorm:"column:defendant_id;type:int(11) unsigned;NOT NULL" json:"defendant_id"`
	OfferId         uint      `gorm:"column:offer_id;type:int(11) unsigned;NOT NULL" json:"offer_id"`
	ComplaintReason string    `gorm:"column:complaint_reason;type:text;NOT NULL" json:"complaint_reason"`
	Time            time.Time `gorm:"column:time;type:timestamp;NOT NULL" json:"time"`
}

func (m *Complaint) TableName() string {
	return "complaint"
}
