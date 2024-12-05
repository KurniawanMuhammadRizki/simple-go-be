package model

type CreateBrandRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdateBrandRequest struct {
	ID   int64   `json:"id" validate:"required"`
	Name *string `json:"name"`
}

type DeleteBrandRequest struct {
	ID int64 `json:"id" validate:"required"`
}

type CreateBrandResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateBrandResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
}

type DeleteBrandResponse struct {
	Message string `json:"message"`
}
