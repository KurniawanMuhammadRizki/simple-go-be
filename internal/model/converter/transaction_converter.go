package converter

import (
	"time"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
)

func ToTransactionEntity(req model.CreateTransactionRequest) entity.Transaction {
	return entity.Transaction{
		CustomerID: req.CustomerID,
		TotalCost:  req.TotalCost,
		CreatedAt:  time.Now(),
	}
}

func ToCreateTransactionResponse(transaction entity.Transaction) model.CreateTransactionResponse {
	return model.CreateTransactionResponse{
		ID:         transaction.ID,
		CustomerID: transaction.CustomerID,
		TotalCost:  transaction.TotalCost,
		CreatedAt:  transaction.CreatedAt.Format(time.RFC3339),
	}
}
