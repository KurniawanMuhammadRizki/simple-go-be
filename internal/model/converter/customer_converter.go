package converter

import (
	"time"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
)

func ToCustomerEntity(req model.CreateCustomerRequest) entity.Customer {
	return entity.Customer{
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func ToCreateCustomerResponse(customer entity.Customer) model.CreateCustomerResponse {
	return model.CreateCustomerResponse{
		ID:        customer.ID,
		Name:      customer.Name,
		CreatedAt: customer.CreatedAt.Format(time.RFC3339),
		UpdatedAt: customer.UpdatedAt.Format(time.RFC3339),
	}
}
