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
