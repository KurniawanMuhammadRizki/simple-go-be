package model

type CreateTransactionDetailRequest struct {
	TransactionID int64 `json:"transaction_id" validate:"required"`
	VoucherID     int64 `json:"voucher_id" validate:"required"`
	Quantity      int   `json:"quantity" validate:"required"`
}

type CreateTransactionDetailResponse struct {
	ID            int64  `json:"id"`
	TransactionID int64  `json:"transaction_id"`
	VoucherID     int64  `json:"voucher_id"`
	Quantity      int    `json:"quantity"`
	SubTotalCost  int64  `json:"sub_total_cost"`
	CreatedAt     string `json:"created_at"`
}
