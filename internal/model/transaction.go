package model

type CreateTransactionRequest struct {
	CustomerID int64 `json:"customer_id" validate:"required"`
	TotalCost  int64 `json:"total_cost" validate:"required"`
}

type CreateTransactionResponse struct {
	ID         int64  `json:"id"`
	CustomerID int64  `json:"customer_id"`
	TotalCost  int64  `json:"total_cost"`
	CreatedAt  string `json:"created_at"`
}

type CreateRedemptionRequest struct {
	CustomerID   int64                   `json:"customer_id" validate:"required"`
	VoucherItems []RedemptionVoucherItem `json:"voucher_items" validate:"required"`
}

type RedemptionVoucherItem struct {
	VoucherID int64 `json:"voucher_id" validate:"required"`
	Quantity  int   `json:"quantity" validate:"required,min=1"`
}

type CreateRedemptionResponse struct {
	TransactionID int64                             `json:"transaction_id"`
	Details       []CreateTransactionDetailResponse `json:"details"`
}
