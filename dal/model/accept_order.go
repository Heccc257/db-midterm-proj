package model

import (
	"database/sql"
)

/*
create table accept_order
(
   order_id             int unsigned not null,
   user_id              int unsigned  not null,
   created_at timestamp NULL DEFAULT NULL,
   status_code          ENUM('canceled', 'in_progress', 'completed') not null default 'in_progress',
   primary key (order_id, user_id)
);
*/

type StatusCode int

const (
	canceled StatusCode = iota
	in_progress
	completed
)

type AcceptOrder struct {
	OrderId    uint         `gorm:"column:order_id;type:int(11) unsigned;primary_key" json:"order_id"`
	UserId     uint         `gorm:"column:user_id;type:int(11) unsigned;NOT NULL" json:"user_id"`
	CreatedAt  sql.NullTime `gorm:"column:created_at;type:timestamp" json:"created_at"`
	StatusCode string       `gorm:"column:status_code;type:enum('canceled','in_progress','completed');default:in_progress;NOT NULL" json:"status_code"`
}

func (m *AcceptOrder) TableName() string {
	return "accept_order"
}
