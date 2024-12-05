package entity

import (
	"time"
)

type Voucher struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name"`
	BrandID     int64     `gorm:"column:brand_id"`
	Brand       Brand     `gorm:"foreignKey:BrandID;references:ID"`
	CostInPoint int64     `gorm:"column:cost_in_points"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamptz;default:now()"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamptz;default:now()"`
}

func (p *Voucher) TableName() string {
	return "vouchers"
}
