package http

import (
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type BrandController struct {
	Usecase usecase.BrandUsecase
	Log     *logrus.Logger
}

func NewBrandController(uc *usecase.BrandUsecase, log *logrus.Logger) *BrandController {
	return &BrandController{
		Usecase: *uc,
		Log:     log,
	}
}

func (p *BrandController) CreateBrand(ctx *fiber.Ctx) error {
	request := new(model.CreateBrandRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}
	resp, err := p.Usecase.CreateBrand(ctx.Context(), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to create brand")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.CreateBrandResponse]{
			Data:    resp,
			Success: false,
			Message: "failed to create brand",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CreateBrandResponse]{
		Data:    resp,
		Success: true,
		Message: "brand created successfully",
	})
}

func (p *BrandController) GetBrandByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid brand id format",
			Success: false,
		})
	}

	brand, err := p.Usecase.GetBrandByID(ctx.Context(), int64(id))
	if err != nil {
		if err.Error() == "brand not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Message: "brand not found",
				Success: false,
			})
		}

		p.Log.WithError(err).Error("failed to get brand")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve brand details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateBrandResponse]{
		Data:    brand,
		Success: true,
		Message: "brand retrieved successfully",
	})
}

func (p *BrandController) UpdateBrand(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid brand id format",
			Success: false,
		})
	}

	request := new(model.UpdateBrandRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}

	brand, err := p.Usecase.UpdateBrand(ctx.Context(), request, int64(id))
	if err != nil {
		p.Log.WithError(err).Error("failed to update brand")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to update brand details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateBrandResponse]{
		Data:    brand,
		Success: true,
		Message: "brand updated successfully",
	})
}

func (p *BrandController) DeleteBrand(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid brand id format",
			Success: false,
		})
	}

	resp, err := p.Usecase.DeleteBrand(ctx.Context(), int64(id))
	if err != nil {
		p.Log.WithError(err).Error("failed to delete brand")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to delete brand",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.DeleteBrandResponse]{
		Data:    resp,
		Success: true,
		Message: "brand deleted successfully",
	})
}
