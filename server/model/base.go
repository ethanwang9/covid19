package model

import "time"

// name: 数据库公共字段
// author: Ethan.Wang
// desc:

type Base struct {
	CreatedAt time.Time `gorm:"column:create_at" json:"create_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
