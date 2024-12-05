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
	// setup use cases
	brandUseCase := usecase.NewBrandUsecase(brandRepository, cfg.Log, cfg.DB)
	// setup controller
	brandController := http.NewBrandController(&brandUseCase, cfg.Log)
	// setup middleware
	authMiddleware := middleware.NewAuth()
	routeConfig := http.Router{
		App:             cfg.App,
		BrandController: brandController,
		AuthMiddleware:  authMiddleware,
	}
	routeConfig.Setup()
}
