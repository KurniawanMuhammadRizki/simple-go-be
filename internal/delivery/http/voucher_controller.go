package http

import (
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/model"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type VoucherController struct {
	Usecase usecase.VoucherUsecase
	Log     *logrus.Logger
}

func NewVoucherController(uc *usecase.VoucherUsecase, log *logrus.Logger) *VoucherController {
	return &VoucherController{
		Usecase: *uc,
		Log:     log,
	}
}

func (p *VoucherController) CreateVoucher(ctx *fiber.Ctx) error {
	request := new(model.CreateVoucherRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}
	resp, err := p.Usecase.CreateVoucher(ctx.Context(), request)
	if err != nil {
		p.Log.WithError(err).Error("failed to create voucher")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[*model.CreateVoucherResponse]{
			Data:    resp,
			Success: false,
			Message: "failed to create voucher",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse[*model.CreateVoucherResponse]{
		Data:    resp,
		Success: true,
		Message: "voucher created successfully",
	})
}

func (p *VoucherController) GetVoucherByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid voucher id format",
			Success: false,
		})
	}

	voucher, err := p.Usecase.GetVoucherByID(ctx.Context(), int64(id))
	if err != nil {
		if err.Error() == "voucher not found" {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
				Message: "voucher not found",
				Success: false,
			})
		}

		p.Log.WithError(err).Error("failed to get voucher")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve voucher details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateVoucherResponse]{
		Data:    voucher,
		Success: true,
		Message: "voucher retrieved successfully",
	})
}

func (p *VoucherController) GetAllByBrand(ctx *fiber.Ctx) error {

	brandID := ctx.QueryInt("id")
	if brandID <= 0 {
		p.Log.Error("invalid or missing brand_id in query")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid or missing brand id",
			Success: false,
		})
	}

	vouchers, err := p.Usecase.GetAllByBrand(ctx.Context(), int64(brandID))
	if err != nil {
		p.Log.WithError(err).Error("failed to get vouchers by brand")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to retrieve vouchers for the specified brand",
			Success: false,
		})
	}

	var voucherPointers []*model.CreateVoucherResponse
	for i := range vouchers {
		voucherPointers = append(voucherPointers, &vouchers[i])
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[[]*model.CreateVoucherResponse]{
		Data:    voucherPointers,
		Success: true,
		Message: "vouchers retrieved successfully",
	})
}

func (p *VoucherController) UpdateVoucher(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid voucher id format",
			Success: false,
		})
	}

	request := new(model.UpdateVoucherRequest)
	if err := ctx.BodyParser(request); err != nil {
		p.Log.WithError(err).Error("failed to parse request body")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid request body format",
			Success: false,
		})
	}

	voucher, err := p.Usecase.UpdateVoucher(ctx.Context(), request, int64(id))
	if err != nil {
		p.Log.WithError(err).Error("failed to update voucher")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to update voucher details",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.CreateVoucherResponse]{
		Data:    voucher,
		Success: true,
		Message: "voucher updated successfully",
	})
}

func (p *VoucherController) DeleteVoucher(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		p.Log.WithError(err).Error("failed to parse id")
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Message: "invalid voucher id format",
			Success: false,
		})
	}

	resp, err := p.Usecase.DeleteVoucher(ctx.Context(), int64(id))
	if err != nil {
		p.Log.WithError(err).Error("failed to delete voucher")
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.WebResponse[any]{
			Message: "failed to delete voucher",
			Success: false,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(model.WebResponse[*model.DeleteVoucherResponse]{
		Data:    resp,
		Success: true,
		Message: "voucher deleted successfully",
	})
}
