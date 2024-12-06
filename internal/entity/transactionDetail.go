package entity

import (
	"time"
)

type TransactionDetail struct {
	ID            int64       `gorm:"column:id;primaryKey;autoIncrement"`
	TransactionID int64       `gorm:"column:transaction_id"`
	VoucherID     int64       `gorm:"column:voucher_id"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	Voucher       Voucher     `gorm:"foreignKey:VoucherID;references:ID"`
	Quantity      int         `gorm:"column:quantity"`
	SubTotalCost  int64       `gorm:"column:sub_total_cost"`
	CreatedAt     time.Time   `gorm:"column:created_at;type:timestamptz;default:now()"`
}

func (p *TransactionDetail) TableName() string {
	return "transaction_details"
}
