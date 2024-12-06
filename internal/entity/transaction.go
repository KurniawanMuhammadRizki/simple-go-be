package entity

import (
	"time"
)

type Transaction struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement"`
	CustomerID int64     `gorm:"column:customer_id"`
	Customer   Customer  `gorm:"foreignKey:CustomerID;references:ID"`
	TotalCost  int64     `gorm:"column:total_cost"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamptz;default:now()"`
}

func (p *Transaction) TableName() string {
	return "transactions"
}
