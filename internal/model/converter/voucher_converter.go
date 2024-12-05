package converter

import (
	"time"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
)

func ToVoucherEntity(req model.CreateVoucherRequest) entity.Voucher {
	return entity.Voucher{
		Name:        req.Name,
		BrandID:     req.BrandID,
		CostInPoint: req.CostInPoint,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ToCreateVoucherResponse(voucher entity.Voucher) model.CreateVoucherResponse {
	return model.CreateVoucherResponse{
		ID:          voucher.ID,
		Name:        voucher.Name,
		BrandID:     voucher.BrandID,
		CostInPoint: voucher.CostInPoint,
		CreatedAt:   voucher.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   voucher.UpdatedAt.Format(time.RFC3339),
	}
}

func ToUpdateVoucherResponse(voucher entity.Voucher) model.UpdateVoucherResponse {
	return model.UpdateVoucherResponse{
		ID:          voucher.ID,
		Name:        voucher.Name,
		BrandID:     voucher.BrandID,
		CostInPoint: voucher.CostInPoint,
		UpdatedAt:   voucher.UpdatedAt.Format(time.RFC3339),
	}
}
