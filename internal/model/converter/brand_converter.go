package converter

import (
	"time"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
)

func ToBrandEntity(req model.CreateBrandRequest) entity.Brand {
	return entity.Brand{
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ToCreateBrandResponse(brand entity.Brand) model.CreateBrandResponse {
	return model.CreateBrandResponse{
		ID:        brand.ID,
		Name:      brand.Name,
		CreatedAt: brand.CreatedAt.Format(time.RFC3339),
		UpdatedAt: brand.UpdatedAt.Format(time.RFC3339),
	}
}

func ToUpdateBrandResponse(brand entity.Brand) model.UpdateBrandResponse {
	return model.UpdateBrandResponse{
		ID:        brand.ID,
		Name:      brand.Name,
		UpdatedAt: brand.UpdatedAt.Format(time.RFC3339),
	}
}
