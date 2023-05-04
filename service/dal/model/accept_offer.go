package model

import (
	"database/sql"
	"time"
)

type AcceptOffer struct {
	OfferId      uint         `gorm:"column:offer_id;type:int(11) unsigned;primary_key" json:"offer_id"`
	UserId       uint         `gorm:"column:user_id;type:int(11) unsigned;NOT NULL" json:"user_id"`
	CreatedAt    sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	CompleteTime time.Time    `gorm:"column:complete_time;type:datetime" json:"complete_time"`
	TaskState    string       `gorm:"column:task_state;type:enum('canceled','in_progress','completed');default:in_progress;NOT NULL" json:"task_state"`
}

func (m *AcceptOffer) TableName() string {
	return "accept_offer"
}
