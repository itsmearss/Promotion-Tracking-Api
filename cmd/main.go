package main

import (
	"submission_promotion_tracking_api/internal/app/repositories"
	"submission_promotion_tracking_api/internal/app/services"
	"submission_promotion_tracking_api/internal/configs"
	"submission_promotion_tracking_api/internal/delivery"

	"github.com/labstack/echo/v4"
)

func main() {

	configs.LoadViperEnv()

	db := configs.InitDatabase()

	e := echo.New()

	PromotionRepo := repositories.NewPromotionRepository(db)

	PromoService := services.NewPromotionService(PromotionRepo)

	delivery.PromotionRoute(e, PromoService)

	e.Logger.Fatal(e.Start(":8080"))
}
