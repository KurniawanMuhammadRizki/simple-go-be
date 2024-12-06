package http

import (
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CustomerController struct {
	Usecase usecase.CustomerUsecase
	Log     *logrus.Logger
}

func NewCustomerController(uc *usecase.CustomerUsecase, log *logrus.Logger) *CustomerController {
	return &CustomerController{
		Usecase: *uc,
		Log:     log,
	}
}

func (p *CustomerController) CreateCustomer(ctx *fiber.Ctx) error {
	request := new(model.CreateCustomerRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}
	resp, err := p.Usecase.CreateCustomer(ctx.Context(), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to create customer")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.CreateCustomerResponse]{
			Data:    resp,
			Success: false,
			Message: "failed to create customer",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CreateCustomerResponse]{
		Data:    resp,
		Success: true,
		Message: "customer created successfully",
	})
}

func (p *CustomerController) GetCustomerByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid customer id format",
			Success: false,
		})
	}

	customer, err := p.Usecase.GetCustomerByID(ctx.Context(), int64(id))
	if err != nil {
		if err.Error() == "customer not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Message: "customer not found",
				Success: false,
			})
		}

		p.Log.WithError(err).Error("failed to get customer")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve customer details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateCustomerResponse]{
		Data:    customer,
		Success: true,
		Message: "customer retrieved successfully",
	})
}
