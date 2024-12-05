package entity

import (
	"time"
)

type Brand struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamptz;default:now()"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamptz;default:now()"`
}

func (p *Brand) TableName() string {
	return "brands"
}
