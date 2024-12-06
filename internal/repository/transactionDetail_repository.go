package repository

import (
	"errors"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionDetailRepository struct {
	Log *logrus.Logger
}

func NewTransactionDetailRepository(log *logrus.Logger) *TransactionDetailRepository {
	return &TransactionDetailRepository{
		Log: log,
	}
}

func (b *TransactionDetailRepository) Save(db *gorm.DB, transactionDetail *entity.TransactionDetail) (*entity.TransactionDetail, error) {
	err := db.Create(transactionDetail).Error
	if err != nil {
		b.Log.Error(err)
		return nil, err
	}

	return transactionDetail, nil
}

// GetByID implements TransactionDetailRepository.
func (b *TransactionDetailRepository) GetByID(db *gorm.DB, id int64) (*entity.TransactionDetail, error) {
	var transactionDetail entity.TransactionDetail
	err := db.Where("id = ?", id).First(&transactionDetail).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("id", id).Error("transactionDetail not found")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}
	return &transactionDetail, nil
}

func (b *TransactionDetailRepository) GetAllByTransaction(db *gorm.DB, transactionID int64) ([]entity.TransactionDetail, error) {
	var transactionDetails []entity.TransactionDetail

	err := db.Where("transaction_id = ?", transactionID).Find(&transactionDetails).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("transaction_id", transactionID).Error("no transactionDetails found for the given transaction")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}

	return transactionDetails, nil
}
