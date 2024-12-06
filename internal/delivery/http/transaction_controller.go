package http

import (
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionController struct {
	Usecase usecase.TransactionUsecase
	Log     *logrus.Logger
}

func NewTransactionController(uc *usecase.TransactionUsecase, log *logrus.Logger) *TransactionController {
	return &TransactionController{
		Usecase: *uc,
		Log:     log,
	}
}

func (p *TransactionController) CreateTransaction(ctx *fiber.Ctx) error {
	request := new(model.CreateTransactionRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}
	resp, err := p.Usecase.CreateTransaction(ctx.Context(), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to create transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.CreateTransactionResponse]{
			Data:    resp,
			Success: false,
			Message: "failed to create transaction",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CreateTransactionResponse]{
		Data:    resp,
		Success: true,
		Message: "transaction created successfully",
	})
}

func (p *TransactionController) GetTransactionByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid transaction id format",
			Success: false,
		})
	}

	transaction, err := p.Usecase.GetTransactionByID(ctx.Context(), int64(id))
	if err != nil {
		if err.Error() == "transaction not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Message: "transaction not found",
				Success: false,
			})
		}

		p.Log.WithError(err).Error("failed to get transaction")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve transaction details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateTransactionResponse]{
		Data:    transaction,
		Success: true,
		Message: "transaction retrieved successfully",
	})
}

func (p *TransactionController) GetAllByCustomer(ctx *fiber.Ctx) error {

	customerID := ctx.QueryInt("id")
	if customerID <= 0 {
		p.Log.Error("invalid or missing customer_id in query")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid or missing customer id",
			Success: false,
		})
	}

	transactions, err := p.Usecase.GetAllByCustomer(ctx.Context(), int64(customerID))
	if err != nil {
		p.Log.WithError(err).Error("failed to get transactions by customer")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve transactions for the specified customer",
			Success: false,
		})
	}

	if len(transactions) == 0 {
		p.Log.Warnf("no transactions found for customer id: %d", customerID)
		return ctx.Status(fiber.StatusNotFound).JSON(model.WebResponse[any]{
			Message: "no transactions found for the specified customer",
			Success: false,
		})
	}

	var transactionPointers []*model.CreateTransactionResponse
	for i := range transactions {
		transactionPointers = append(transactionPointers, &transactions[i])
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[[]*model.CreateTransactionResponse]{
		Data:    transactionPointers,
		Success: true,
		Message: "transactions retrieved successfully",
	})
}
