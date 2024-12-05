package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model/converter"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VoucherUsecase interface {
	CreateVoucher(ctx context.Context, req *model.CreateVoucherRequest) (*model.CreateVoucherResponse, error)
	GetVoucherByID(ctx context.Context, id int64) (*model.CreateVoucherResponse, error)
	UpdateVoucher(ctx context.Context, req *model.UpdateVoucherRequest, id int64) (*model.CreateVoucherResponse, error)
	DeleteVoucher(ctx context.Context, id int64) (*model.DeleteVoucherResponse, error)
}

type voucherUsecase struct {
	VoucherRepository *repository.VoucherRepository
	Log               *logrus.Logger
	DB                *gorm.DB
}

func NewVoucherUsecase(
	voucherRepository *repository.VoucherRepository,
	log *logrus.Logger,
	db *gorm.DB,
) VoucherUsecase {
	return &voucherUsecase{
		VoucherRepository: voucherRepository,
		Log:               log,
		DB:                db,
	}
}

func (p *voucherUsecase) CreateVoucher(ctx context.Context, req *model.CreateVoucherRequest) (*model.CreateVoucherResponse, error) {
	tx := p.DB.Begin()
	voucher := converter.ToVoucherEntity(*req)
	savedVoucher, err := p.VoucherRepository.Save(tx, &voucher)

	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to save voucher")
		return nil, err
	}
	response := converter.ToCreateVoucherResponse(*savedVoucher)
	return &response, tx.Commit().Error
}

func (p *voucherUsecase) GetVoucherByID(ctx context.Context, id int64) (*model.CreateVoucherResponse, error) {
	voucher, err := p.VoucherRepository.GetByID(p.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("voucher not found")
		}
		p.Log.WithError(err).Error("failed to get voucher")
		return nil, err
	}

	response := converter.ToCreateVoucherResponse(*voucher)
	return &response, nil
}

func (p *voucherUsecase) GetAllByBrand(ctx context.Context, brandID int64) ([]model.CreateVoucherResponse, error) {
	vouchers, err := p.VoucherRepository.GetAllByBrand(p.DB, brandID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no vouchers found for the given brand")
		}
		p.Log.WithError(err).Error("failed to get vouchers by brand")
		return nil, err
	}

	// Convert vouchers to response
	var responses []model.CreateVoucherResponse
	for _, voucher := range vouchers {
		responses = append(responses, converter.ToCreateVoucherResponse(voucher))
	}

	return responses, nil
}

func (p *voucherUsecase) UpdateVoucher(ctx context.Context, req *model.UpdateVoucherRequest, id int64) (*model.CreateVoucherResponse, error) {
	tx := p.DB.Begin()

	existingVoucher, err := p.VoucherRepository.GetByID(tx, id)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("voucher not found")
		}
		p.Log.WithError(err).Error("failed to get voucher")
		return nil, err
	}

	if req.Name != nil {
		existingVoucher.Name = *req.Name
	}

	existingVoucher.UpdatedAt = time.Now()

	updatedVoucher, err := p.VoucherRepository.Update(tx, existingVoucher)
	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to update voucher")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	response := converter.ToCreateVoucherResponse(*updatedVoucher)
	return &response, nil
}

func (p *voucherUsecase) DeleteVoucher(ctx context.Context, id int64) (*model.DeleteVoucherResponse, error) {
	tx := p.DB.Begin()

	_, err := p.VoucherRepository.GetByID(tx, id)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("voucher not found")
		}
		p.Log.WithError(err).Error("failed to get voucher")
		return nil, err
	}

	err = p.VoucherRepository.Delete(tx, id)
	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to delete voucher")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return &model.DeleteVoucherResponse{
		Message: "voucher deleted successfully",
	}, nil
}
