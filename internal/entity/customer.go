package entity

import (
	"time"
)

type Customer struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name         string    `gorm:"column:name"`
	PointBalance int64     `gorm:"column:point_balance"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamptz;default:now()"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamptz;default:now()"`
}

func (p *Customer) TableName() string {
	return "customers"
}
