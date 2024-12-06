package usecase

import (
	"context"
	"errors"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model/converter"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/repository"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionDetailUsecase interface {
	CreateTransactionDetail(ctx context.Context, req *model.CreateTransactionDetailRequest) (*model.CreateTransactionDetailResponse, error)
	GetTransactionDetailByID(ctx context.Context, id int64) (*model.CreateTransactionDetailResponse, error)
	GetAllByTransaction(ctx context.Context, transactionID int64) ([]model.CreateTransactionDetailResponse, error)
}

type transactionDetailUsecase struct {
	TransactionDetailRepository *repository.TransactionDetailRepository
	Log                         *logrus.Logger
	DB                          *gorm.DB
}

func NewTransactionDetailUsecase(
	transactionDetailRepository *repository.TransactionDetailRepository,
	log *logrus.Logger,
	db *gorm.DB,
) TransactionDetailUsecase {
	return &transactionDetailUsecase{
		TransactionDetailRepository: transactionDetailRepository,
		Log:                         log,
		DB:                          db,
	}
}

func (p *transactionDetailUsecase) CreateTransactionDetail(ctx context.Context, req *model.CreateTransactionDetailRequest) (*model.CreateTransactionDetailResponse, error) {
	tx := p.DB.Begin()
	transactionDetail := converter.ToTransactionDetailEntity(*req)
	savedTransactionDetail, err := p.TransactionDetailRepository.Save(tx, &transactionDetail)

	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to save transactionDetail")
		return nil, err
	}
	response := converter.ToCreateTransactionDetailResponse(*savedTransactionDetail)
	return &response, tx.Commit().Error
}

func (p *transactionDetailUsecase) GetTransactionDetailByID(ctx context.Context, id int64) (*model.CreateTransactionDetailResponse, error) {
	transactionDetail, err := p.TransactionDetailRepository.GetByID(p.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transactionDetail not found")
		}
		p.Log.WithError(err).Error("failed to get transactionDetail")
		return nil, err
	}

	response := converter.ToCreateTransactionDetailResponse(*transactionDetail)
	return &response, nil
}

func (p *transactionDetailUsecase) GetAllByTransaction(ctx context.Context, transactionID int64) ([]model.CreateTransactionDetailResponse, error) {
	transactionDetails, err := p.TransactionDetailRepository.GetAllByTransaction(p.DB, transactionID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no transactionDetails found for the given transaction")
		}
		p.Log.WithError(err).Error("failed to get transactionDetails by transaction")
		return nil, err
	}

	// Convert transactionDetails to response
	var responses []model.CreateTransactionDetailResponse
	for _, transactionDetail := range transactionDetails {
		responses = append(responses, converter.ToCreateTransactionDetailResponse(transactionDetail))
	}

	return responses, nil
}
