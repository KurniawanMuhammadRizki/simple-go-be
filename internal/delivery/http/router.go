package http

import "github.com/gofiber/fiber/v2"

type Router struct {
	App                *fiber.App
	BrandController    *BrandController
	VoucherController  *VoucherController
	CustomerController *CustomerController
	AuthMiddleware     fiber.Handler
}

type router interface {
	Setup()
	registerPublicEndpoints()
	registerPrivateEndpoints()
}

func NewRouter(app *fiber.App, brandController *BrandController, voucherController *VoucherController, customerController *CustomerController, authMiddleware fiber.Handler) router {
	return &Router{
		App:                app,
		BrandController:    brandController,
		VoucherController:  voucherController,
		CustomerController: customerController,
		AuthMiddleware:     authMiddleware,
	}
}

// Setup implements router.
func (r *Router) Setup() {
	r.registerPublicEndpoints()
	r.registerPrivateEndpoints()
}

// registerPrivateEndpoints implements router.
func (r *Router) registerPrivateEndpoints() {
	r.App.Use(r.AuthMiddleware)
	r.App.Get("/secrit", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "This is a secret endpoint",
		})
	})
}

// registerPublicEndpoints implements router.
func (r *Router) registerPublicEndpoints() {
	//brand
	r.App.Post("/brands", r.BrandController.CreateBrand)
	r.App.Get("/brands/:id", r.BrandController.GetBrandByID)
	r.App.Put("/brands/:id", r.BrandController.UpdateBrand)
	r.App.Delete("/brands/:id", r.BrandController.DeleteBrand)
	//voucher
	r.App.Post("/vouchers", r.VoucherController.CreateVoucher)
	r.App.Get("/vouchers/brand", r.VoucherController.GetAllByBrand)
	r.App.Get("/vouchers/:id", r.VoucherController.GetVoucherByID)
	r.App.Put("/vouchers/:id", r.VoucherController.UpdateVoucher)
	r.App.Delete("/vouchers/:id", r.VoucherController.DeleteVoucher)
	//customer
	r.App.Post("/customers", r.CustomerController.CreateCustomer)
	r.App.Get("/customers/:id", r.CustomerController.GetCustomerByID)

}
