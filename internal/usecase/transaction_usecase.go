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
	CreateRedemption(ctx context.Context, req *model.CreateRedemptionRequest) (*model.CreateRedemptionResponse, error)
}

type transactionUsecase struct {
	TransactionRepository    *repository.TransactionRepository
	VoucherRepository        *repository.VoucherRepository
	TransactionDetailUsecase TransactionDetailUsecase
	Log                      *logrus.Logger
	DB                       *gorm.DB
}

func NewTransactionUsecase(
	transactionRepository *repository.TransactionRepository,
	transactionDetailUsecase TransactionDetailUsecase,
	log *logrus.Logger,
	db *gorm.DB,
) TransactionUsecase {
	return &transactionUsecase{
		TransactionRepository:    transactionRepository,
		TransactionDetailUsecase: transactionDetailUsecase,
		Log:                      log,
		DB:                       db,
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

func (p *transactionUsecase) CreateRedemption(ctx context.Context, req *model.CreateRedemptionRequest) (*model.CreateRedemptionResponse, error) {
	tx := p.DB.Begin()

	if p.TransactionDetailUsecase == nil {
		p.Log.Error("TransactionDetailUsecase is nil")
		tx.Rollback()
		return nil, errors.New("TransactionDetailUsecase is not initialized")
	}

	p.Log.Info("Starting redemption creation process")

	transactionReq := &model.CreateTransactionRequest{
		CustomerID: req.CustomerID,
		TotalCost:  0, // dihitung nanti
	}

	transaction, err := p.CreateTransaction(ctx, transactionReq)
	if err != nil {
		p.Log.WithError(err).Error("Failed to create transaction")
		tx.Rollback()
		return nil, err
	}

	p.Log.Infof("Transaction created successfully: ID=%d", transaction.ID)

	var totalCost int64
	var redemptionDetails []model.CreateTransactionDetailResponse

	for _, item := range req.VoucherItems {
		transactionDetailReq := &model.CreateTransactionDetailRequest{
			TransactionID: transaction.ID,
			VoucherID:     item.VoucherID,
			Quantity:      item.Quantity,
		}

		if p.TransactionDetailUsecase == nil {
			p.Log.Error("TransactionDetailUsecase is nil")
			tx.Rollback()
			return nil, errors.New("TransactionDetailUsecase is not initialized")
		}

		detailTransaction, err := p.TransactionDetailUsecase.CreateTransactionDetail(ctx, transactionDetailReq)
		if err != nil {
			p.Log.WithError(err).Errorf("Failed to create transaction detail for VoucherID=%d", item.VoucherID)
			tx.Rollback()
			return nil, err
		}

		p.Log.Infof("Transaction detail created: TransactionID=%d, VoucherID=%d, SubTotalCost=%d",
			detailTransaction.TransactionID, detailTransaction.VoucherID, detailTransaction.SubTotalCost)

		totalCost += detailTransaction.SubTotalCost
		redemptionDetails = append(redemptionDetails, *detailTransaction)
	}

	transaction.TotalCost = totalCost
	existingTransaction, err := p.TransactionRepository.GetByID(tx, transaction.ID)
	if err != nil {
		p.Log.WithError(err).Error("Transaction not foundd")
		tx.Rollback()
		return nil, err
	}
	existingTransaction.TotalCost = totalCost
	p.Log.Infof("Transaction Entity to Update: %+v", existingTransaction)

	_, err = p.TransactionRepository.Update(tx, existingTransaction)
	if err != nil {
		p.Log.WithError(err).Error("Failed to update transaction")
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		p.Log.WithError(err).Error("Failed to commit transaction")
		return nil, err
	}

	p.Log.Info("Redemption creation process completed successfully")
	return &model.CreateRedemptionResponse{
		TransactionID: transaction.ID,
		Details:       redemptionDetails,
	}, nil
}
