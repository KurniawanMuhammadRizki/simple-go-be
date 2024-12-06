package http

import (
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionDetailController struct {
	Usecase usecase.TransactionDetailUsecase
	Log     *logrus.Logger
}

func NewTransactionDetailController(uc *usecase.TransactionDetailUsecase, log *logrus.Logger) *TransactionDetailController {
	return &TransactionDetailController{
		Usecase: *uc,
		Log:     log,
	}
}

func (p *TransactionDetailController) CreateTransactionDetail(ctx *fiber.Ctx) error {
	request := new(model.CreateTransactionDetailRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}
	resp, err := p.Usecase.CreateTransactionDetail(ctx.Context(), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to create transaction detail")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.CreateTransactionDetailResponse]{
			Data:    resp,
			Success: false,
			Message: "failed to create transaction detail",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CreateTransactionDetailResponse]{
		Data:    resp,
		Success: true,
		Message: "transaction detail created successfully",
	})
}

func (p *TransactionDetailController) GetTransactionDetailByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid transaction detail id format",
			Success: false,
		})
	}

	transactionDetail, err := p.Usecase.GetTransactionDetailByID(ctx.Context(), int64(id))
	if err != nil {
		if err.Error() == "transaction detail not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Message: "transaction detail not found",
				Success: false,
			})
		}

		p.Log.WithError(err).Error("failed to get transaction detail")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve transaction details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateTransactionDetailResponse]{
		Data:    transactionDetail,
		Success: true,
		Message: "transaction detail retrieved successfully",
	})
}

func (p *TransactionDetailController) GetAllByTransaction(ctx *fiber.Ctx) error {

	transactionID := ctx.QueryInt("id")
	if transactionID <= 0 {
		p.Log.Error("invalid or missing transaction_id in query")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid or missing transaction id",
			Success: false,
		})
	}

	transactionDetails, err := p.Usecase.GetAllByTransaction(ctx.Context(), int64(transactionID))
	if err != nil {
		p.Log.WithError(err).Error("failed to get transactionDetails by transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve transactionDetails for the specified transaction",
			Success: false,
		})
	}

	var transactionDetailPointers []*model.CreateTransactionDetailResponse
	for i := range transactionDetails {
		transactionDetailPointers = append(transactionDetailPointers, &transactionDetails[i])
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[[]*model.CreateTransactionDetailResponse]{
		Data:    transactionDetailPointers,
		Success: true,
		Message: "transactionDetails retrieved successfully",
	})
}
