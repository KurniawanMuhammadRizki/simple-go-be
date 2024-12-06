package repository

import (
	"errors"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	Log *logrus.Logger
}

func NewTransactionRepository(log *logrus.Logger) *TransactionRepository {
	return &TransactionRepository{
		Log: log,
	}
}

func (b *TransactionRepository) Save(db *gorm.DB, transaction *entity.Transaction) (*entity.Transaction, error) {
	err := db.Create(transaction).Error
	if err != nil {
		b.Log.Error(err)
		return nil, err
	}

	return transaction, nil
}

// GetByID implements TransactionRepository.
func (b *TransactionRepository) GetByID(db *gorm.DB, id int64) (*entity.Transaction, error) {
	var transaction entity.Transaction
	err := db.Where("id = ?", id).First(&transaction).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("id", id).Error("transaction not found")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}
	return &transaction, nil
}

func (b *TransactionRepository) GetAllByCustomer(db *gorm.DB, CustomerID int64) ([]entity.Transaction, error) {
	var transactions []entity.Transaction

	err := db.Where("customer_id = ?", CustomerID).Find(&transactions).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("customer_id", CustomerID).Error("no transactions found for the given customer")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}

	return transactions, nil
}

func (b *TransactionRepository) Update(db *gorm.DB, transaction *entity.Transaction) (*entity.Transaction, error) {
	err := db.Where("id = ?", transaction.ID).Updates(transaction).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("id", transaction.ID).Error("transaction not found")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}

	return b.GetByID(db, transaction.ID)
}
