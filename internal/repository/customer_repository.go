package repository

import (
	"errors"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	Log *logrus.Logger
}

func NewCustomerRepository(log *logrus.Logger) *CustomerRepository {
	return &CustomerRepository{
		Log: log,
	}
}

func (b *CustomerRepository) Save(db *gorm.DB, customer *entity.Customer) (*entity.Customer, error) {
	err := db.Create(customer).Error
	if err != nil {
		b.Log.Error(err)
		return nil, err
	}

	return customer, nil
}

// GetByID implements CustomerRepository.
func (b *CustomerRepository) GetByID(db *gorm.DB, id int64) (*entity.Customer, error) {
	var customer entity.Customer
	err := db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			b.Log.WithField("id", id).Error("customer not found")
			return nil, err
		}
		b.Log.Error(err)
		return nil, err
	}
	return &customer, nil
}
