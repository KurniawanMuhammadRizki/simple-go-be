package config

import (
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/delivery/http"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/delivery/http/middleware"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/repository"
	"github.com/KurniawanMuhammadRizki/simple-go-be/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB     *gorm.DB
	App    *fiber.App
	Log    *logrus.Logger
	Config *viper.Viper
}

func (cfg *AppConfig) Run() {
	// setup repositories
	brandRepository := repository.NewBrandRepository(cfg.Log)
	voucherRepository := repository.NewVoucherRepository(cfg.Log)
	customerRepository := repository.NewCustomerRepository(cfg.Log)
	transactionRepository := repository.NewTransactionRepository(cfg.Log)
	// setup use cases
	brandUseCase := usecase.NewBrandUsecase(brandRepository, cfg.Log, cfg.DB)
	voucherUseCase := usecase.NewVoucherUsecase(voucherRepository, cfg.Log, cfg.DB)
	customerUseCase := usecase.NewCustomerUsecase(customerRepository, cfg.Log, cfg.DB)
	transactionUseCase := usecase.NewTransactionUsecase(transactionRepository, cfg.Log, cfg.DB)
	// setup controller
	brandController := http.NewBrandController(&brandUseCase, cfg.Log)
	voucherController := http.NewVoucherController(&voucherUseCase, cfg.Log)
	customerController := http.NewCustomerController(&customerUseCase, cfg.Log)
	transactionController := http.NewTransactionController(&transactionUseCase, cfg.Log)
	// setup middleware
	authMiddleware := middleware.NewAuth()
	routeConfig := http.Router{
		App:                   cfg.App,
		BrandController:       brandController,
		VoucherController:     voucherController,
		CustomerController:    customerController,
		TransactionController: transactionController,
		AuthMiddleware:        authMiddleware,
	}
	routeConfig.Setup()
}
