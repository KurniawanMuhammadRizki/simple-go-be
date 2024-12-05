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

type BrandUsecase interface {
	CreateBrand(ctx context.Context, req *model.CreateBrandRequest) (*model.CreateBrandResponse, error)
	GetBrandByID(ctx context.Context, id int64) (*model.CreateBrandResponse, error)
	UpdateBrand(ctx context.Context, req *model.UpdateBrandRequest, id int64) (*model.CreateBrandResponse, error)
	DeleteBrand(ctx context.Context, id int64) (*model.DeleteBrandResponse, error)
}

type brandUsecase struct {
	BrandRepository *repository.BrandRepository
	Log             *logrus.Logger
	DB              *gorm.DB
}

func NewBrandUsecase(
	brandRepository *repository.BrandRepository,
	log *logrus.Logger,
	db *gorm.DB,
) BrandUsecase {
	return &brandUsecase{
		BrandRepository: brandRepository,
		Log:             log,
		DB:              db,
	}
}

func (p *brandUsecase) CreateBrand(ctx context.Context, req *model.CreateBrandRequest) (*model.CreateBrandResponse, error) {
	tx := p.DB.Begin()
	brand := converter.ToBrandEntity(*req)
	savedBrand, err := p.BrandRepository.Save(tx, &brand)

	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to save brand")
		return nil, err
	}
	response := converter.ToCreateBrandResponse(*savedBrand)
	return &response, tx.Commit().Error
}

func (p *brandUsecase) GetBrandByID(ctx context.Context, id int64) (*model.CreateBrandResponse, error) {
	brand, err := p.BrandRepository.GetByID(p.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("brand not found")
		}
		p.Log.WithError(err).Error("failed to get brand")
		return nil, err
	}

	response := converter.ToCreateBrandResponse(*brand)
	return &response, nil
}

func (p *brandUsecase) UpdateBrand(ctx context.Context, req *model.UpdateBrandRequest, id int64) (*model.CreateBrandResponse, error) {
	tx := p.DB.Begin()

	existingBrand, err := p.BrandRepository.GetByID(tx, id)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("brand not found")
		}
		p.Log.WithError(err).Error("failed to get brand")
		return nil, err
	}

	if req.Name != nil {
		existingBrand.Name = *req.Name
	}

	existingBrand.UpdatedAt = time.Now()

	updatedBrand, err := p.BrandRepository.Update(tx, existingBrand)
	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to update brand")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	response := converter.ToCreateBrandResponse(*updatedBrand)
	return &response, nil
}

func (p *brandUsecase) DeleteBrand(ctx context.Context, id int64) (*model.DeleteBrandResponse, error) {
	tx := p.DB.Begin()

	_, err := p.BrandRepository.GetByID(tx, id)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("brand not found")
		}
		p.Log.WithError(err).Error("failed to get brand")
		return nil, err
	}

	err = p.BrandRepository.Delete(tx, id)
	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to delete brand")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return &model.DeleteBrandResponse{
		Message: "brand deleted successfully",
	}, nil
}
