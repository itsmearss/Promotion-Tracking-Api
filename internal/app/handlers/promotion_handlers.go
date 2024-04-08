package handlers

import (
	"net/http"
	"submission_promotion_tracking_api/internal/app/models"
	"submission_promotion_tracking_api/internal/app/services"
	"submission_promotion_tracking_api/utils/exception"

	"github.com/labstack/echo/v4"
)

func PSQLCreatePromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	// Return an Echo handler function for creating a new promotion
	return func(c echo.Context) error {
		// Declare a variable to store promotion data from the request
		var promo models.Promotion

		// Bind promotion data from the request body to the promo struct
		if err := c.Bind(&promo); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
		}

		// Call the PromotionService to create a new promotion
		createdPromo, err := PromoService.CreatePromotion(promo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create promotion")
		}

		// Return a JSON response with the newly created promotion
		return c.JSON(http.StatusCreated, createdPromo)
	}
}

func PSQLGetAllPromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve all promotions from the database using the PromotionService
		promotion, err := PromoService.GetAllPromotions()
		if err != nil {
			// Return HTTP error with status InternalServerError (500) if failed to retrieve promotions
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve promotions: "+err.Error())
		}
		// Return JSON response with retrieved promotions and HTTP status OK (200)
		return c.JSON(http.StatusOK, promotion)
	}
}

func PSQLGetPromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the promotion ID from the request parameters
		promotionID := c.Param("promotion_id")

		// Retrieve the promotion details from the database using the PromotionService
		promo, err := PromoService.GetPromotionbyPromotionID(promotionID)
		if err != nil {
			// If the promotion is not found, return an HTTP error with status NotFound (404)
			if e, ok := err.(*exception.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			// Return HTTP error with status InternalServerError (500) if failed to get promotion
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}
		// Return JSON response with the retrieved promotion and HTTP status OK (200)
		return c.JSON(http.StatusOK, promo)
	}
}

func PSQLUpdatePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the promotion ID from the request parameters
		promotionID := c.Param("promotion_id")

		// Retrieve the promotion details from the database using the PromotionService
		promo, err := PromoService.GetPromotionbyPromotionID(promotionID)
		if err != nil {
			// If the promotion is not found, return an HTTP error with status NotFound (404)
			if e, ok := err.(*exception.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			// Return HTTP error with status InternalServerError (500) if failed to get promotion
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get promotion")
		}

		// Bind the request body to the promotion struct
		if err := c.Bind(&promo); err != nil {
			// If the promotion data provided in the request body is invalid, return an HTTP error with status BadRequest (400)
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid promotion data")
		}

		// Update the promotion using the updated promotion data
		updatedPromo, err := PromoService.UpdatePromotionbyPromotionID(promo)
		if err != nil {
			// Return HTTP error with status InternalServerError (500) if failed to update promotion
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update promotion")
		}

		// Return JSON response with the updated promotion and HTTP status OK (200)
		return c.JSON(http.StatusOK, updatedPromo)
	}
}

func PSQLDeletePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract the promotion ID from the request parameters
		promotionID := c.Param("promotion_id")

		// Attempt to delete the promotion by its ID using the PromotionService
		if err := PromoService.DeletePromotionbyPromotionID(promotionID); err != nil {
			// If the promotion is not found, return an HTTP error with status NotFound (404)
			if e, ok := err.(*exception.NotFoundError); ok {
				return echo.NewHTTPError(http.StatusNotFound, e.Error())
			}
			// Return HTTP error with status InternalServerError (500) if failed to delete promotion
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete promotion")
		}
		// Return HTTP status NoContent (204) indicating successful deletion of promotion
		return c.JSON(http.StatusNoContent, "Promotion Data deleted successfully")
	}
}
