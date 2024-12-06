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

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (*model.CreateTransactionResponse, error)
	GetTransactionByID(ctx context.Context, id int64) (*model.CreateTransactionResponse, error)
	GetAllByCustomer(ctx context.Context, customerID int64) ([]model.CreateTransactionResponse, error)
}

type transactionUsecase struct {
	TransactionRepository *repository.TransactionRepository
	Log                   *logrus.Logger
	DB                    *gorm.DB
}

func NewTransactionUsecase(
	transactionRepository *repository.TransactionRepository,
	log *logrus.Logger,
	db *gorm.DB,
) TransactionUsecase {
	return &transactionUsecase{
		TransactionRepository: transactionRepository,
		Log:                   log,
		DB:                    db,
	}
}

func (p *transactionUsecase) CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (*model.CreateTransactionResponse, error) {
	tx := p.DB.Begin()
	transaction := converter.ToTransactionEntity(*req)
	savedTransaction, err := p.TransactionRepository.Save(tx, &transaction)

	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to save transaction")
		return nil, err
	}
	response := converter.ToCreateTransactionResponse(*savedTransaction)
	return &response, tx.Commit().Error
}

func (p *transactionUsecase) GetTransactionByID(ctx context.Context, id int64) (*model.CreateTransactionResponse, error) {
	transaction, err := p.TransactionRepository.GetByID(p.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		p.Log.WithError(err).Error("failed to get transaction")
		return nil, err
	}

	response := converter.ToCreateTransactionResponse(*transaction)
	return &response, nil
}

func (p *transactionUsecase) GetAllByCustomer(ctx context.Context, customerID int64) ([]model.CreateTransactionResponse, error) {
	transactions, err := p.TransactionRepository.GetAllByCustomer(p.DB, customerID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no transactions found for the given customer")
		}
		p.Log.WithError(err).Error("failed to get transactions by customer")
		return nil, err
	}

	var responses []model.CreateTransactionResponse
	for _, transaction := range transactions {
		responses = append(responses, converter.ToCreateTransactionResponse(transaction))
	}

	return responses, nil
}
