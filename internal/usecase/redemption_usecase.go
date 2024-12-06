package usecase

import (
	"context"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RedemptionUsecase interface {
	CreateRedemption(ctx context.Context, req *model.CreateRedemptionRequest) (*model.CreateRedemptionResponse, error)
}

type redemptionUsecase struct {
	TransactionUsecase       TransactionUsecase
	TransactionDetailUsecase TransactionDetailUsecase
	VoucherRepository        *repository.VoucherRepository
	DB                       *gorm.DB
	Log                      *logrus.Logger
}

// func (r *redemptionUsecase) CreateRedemption(ctx context.Context, req *model.CreateRedemptionRequest) (*model.CreateRedemptionResponse, error) {
// 	tx := r.DB.Begin()
// 	defer func() {
// 		if r := recover(); r != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// Buat transaksi utama
// 	transactionReq := &model.CreateTransactionRequest{
// 		CustomerID: req.CustomerID,
// 		TotalCost:  0, // akan dihitung nanti
// 	}
// 	transaction, err := r.TransactionUsecase.CreateTransaction(ctx, transactionReq)
// 	if err != nil {
// 		tx.Rollback()
// 		return nil, err
// 	}

// 	// Total biaya akan dihitung dari voucher
// 	var totalCost int64
// 	var redemptionDetails []model.CreateTransactionDetailResponse

// 	// Buat detail transaksi untuk setiap voucher
// 	for _, item := range req.VoucherItems {
// 		voucher, err := r.VoucherRepository.GetByID(tx, item.VoucherID)
// 		if err != nil {
// 			tx.Rollback()
// 			return nil, err
// 		}

// 		detailReq := &model.CreateTransactionDetailRequest{
// 			TransactionID: transaction.ID,
// 			VoucherID:     item.VoucherID,
// 			Quantity:      item.Quantity,
// 			CostInPoint:   voucher.CostInPoint,
// 		}

// 		detail, err := r.TransactionDetailUsecase.CreateTransactionDetail(ctx, detailReq)
// 		if err != nil {
// 			tx.Rollback()
// 			return nil, err
// 		}

// 		totalCost += detail.SubTotalCost
// 		redemptionDetails = append(redemptionDetails, *detail)
// 	}

// 	// Update total cost transaksi
// 	transaction.TotalCost = totalCost

// 	// Commit transaksi
// 	if err := tx.Commit().Error; err != nil {
// 		return nil, err
// 	}

// 	return &model.CreateRedemptionResponse{
// 		TransactionID: transaction.ID,
// 		Details:       redemptionDetails,
// 	}, nil
// }
