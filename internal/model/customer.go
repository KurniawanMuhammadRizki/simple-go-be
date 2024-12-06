package model

type CreateCustomerRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateCustomerResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
