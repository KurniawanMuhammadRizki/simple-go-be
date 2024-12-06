package converter

import (
	"time"

	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/entity"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
)

func ToTransactionDetailEntity(req model.CreateTransactionDetailRequest) entity.TransactionDetail {
	return entity.TransactionDetail{
		TransactionID: req.TransactionID,
		VoucherID:     req.VoucherID,
		Quantity:      int(req.Quantity),
		CreatedAt:     time.Now(),
	}
}

func ToCreateTransactionDetailResponse(transactionDetail entity.TransactionDetail) model.CreateTransactionDetailResponse {
	return model.CreateTransactionDetailResponse{
		ID:            transactionDetail.ID,
		TransactionID: transactionDetail.TransactionID,
		VoucherID:     transactionDetail.VoucherID,
		Quantity:      int(transactionDetail.Quantity),
		SubTotalCost:  int64(transactionDetail.SubTotalCost),
		CreatedAt:     transactionDetail.CreatedAt.Format(time.RFC3339),
	}
}
