package model

import (
	"database/sql"
)

type StatusCode int

const (
	canceled StatusCode = iota
	in_progress
	completed
)

type AcceptOffer struct {
	OfferId    uint         `gorm:"column:offer_id;type:int(11) unsigned;primary_key" json:"offer_id"`
	UserId     uint         `gorm:"column:user_id;type:int(11) unsigned;NOT NULL" json:"user_id"`
	CreatedAt  sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	StatusCode string       `gorm:"column:status_code;type:enum('canceled','in_progress','completed');default:in_progress;NOT NULL" json:"status_code"`
}

func (m *AcceptOffer) TableName() string {
	return "accept_offer"
}
