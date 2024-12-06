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

type CustomerUsecase interface {
	CreateCustomer(ctx context.Context, req *model.CreateCustomerRequest) (*model.CreateCustomerResponse, error)
	GetCustomerByID(ctx context.Context, id int64) (*model.CreateCustomerResponse, error)
}

type customerUsecase struct {
	CustomerRepository *repository.CustomerRepository
	Log                *logrus.Logger
	DB                 *gorm.DB
}

func NewCustomerUsecase(
	customerRepository *repository.CustomerRepository,
	log *logrus.Logger,
	db *gorm.DB,
) CustomerUsecase {
	return &customerUsecase{
		CustomerRepository: customerRepository,
		Log:                log,
		DB:                 db,
	}
}

func (p *customerUsecase) CreateCustomer(ctx context.Context, req *model.CreateCustomerRequest) (*model.CreateCustomerResponse, error) {
	tx := p.DB.Begin()
	customer := converter.ToCustomerEntity(*req)
	savedCustomer, err := p.CustomerRepository.Save(tx, &customer)

	if err != nil {
		tx.Rollback()
		p.Log.WithError(err).Error("failed to save customer")
		return nil, err
	}
	response := converter.ToCreateCustomerResponse(*savedCustomer)
	return &response, tx.Commit().Error
}

func (p *customerUsecase) GetCustomerByID(ctx context.Context, id int64) (*model.CreateCustomerResponse, error) {
	customer, err := p.CustomerRepository.GetByID(p.DB, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("customer not found")
		}
		p.Log.WithError(err).Error("failed to get customer")
		return nil, err
	}

	response := converter.ToCreateCustomerResponse(*customer)
	return &response, nil
}
