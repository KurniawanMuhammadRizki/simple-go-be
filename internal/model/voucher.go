package model

type CreateVoucherRequest struct {
	Name        string `json:"name" validate:"required"`
	BrandId     int64  `json:"brand_id" validate:"required"`
	CostInPoint int64  `json:"cost_in_point" validate:"required"`
}

type UpdateVoucherRequest struct {
	ID          int64   `json:"id" validate:"required"`
	Name        *string `json:"name"`
	BrandId     int64   `json:"brand_id" validate:"required"`
	CostInPoint int64   `json:"cost_in_point" validate:"required"`
}

type DeleteVoucherRequest struct {
	ID int64 `json:"id" validate:"required"`
}

type CreateVoucherResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	BrandId     int64  `json:"brand_id"`
	CostInPoint int64  `json:"cost_in_point"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateVoucherResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	BrandId     int64  `json:"brand_id"`
	CostInPoint int64  `json:"cost_in_point"`
	UpdatedAt   string `json:"updated_at"`
}

type DeleteVoucherResponse struct {
	Message string `json:"message"`
}
