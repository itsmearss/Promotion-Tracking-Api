package services

import (
	"errors"
	"submission_promotion_tracking_api/internal/app/models"
	"submission_promotion_tracking_api/internal/app/repositories"
	"submission_promotion_tracking_api/utils/exception"

	"gorm.io/gorm"
)

// PromotionService provides promotion-related services
type PromotionService interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}

type PromotionServiceImpl struct {
	PromotionRepo repositories.PromotionRepository
}

// NewPromotionService creates a new instance of PromotionService
func NewPromotionService(PromotionRepo repositories.PromotionRepository) *PromotionServiceImpl {
	return &PromotionServiceImpl{
		PromotionRepo: PromotionRepo,
	}
}

// CreatePromotion creates a new promotion
func (s *PromotionServiceImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	// Call the repository method to create a new promotion
	return s.PromotionRepo.CreatePromotion(promo)
}

// GetAllPromotions that already recorded on database
func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error) {
	return s.PromotionRepo.GetAllPromotions()
}

// GetPromotionByPromotionID will throw data based on promotionID request
func (s *PromotionServiceImpl) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	// Delegate the retrieval operation to the underlying promotion repository (s.PromotionRepo)
	promo, err := s.PromotionRepo.GetPromotionbyPromotionID(promotionID)
	if err != nil {
		// If the promotion is not found, return a custom error indicating that the promotion was not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exception.PromotionIDNotFoundError{
				Message:     "Promotion Not Found",
				PromotionID: promotionID,
			}
		}
		// If any other error occurs during the retrieval operation, return the encountered error
		return models.Promotion{}, err
	}
	// Return the retrieved promotion along with a nil error if the operation is successful
	return promo, nil
}

// UpdatePromotion will update data based on promotionID request
func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	// Perform duplicate check and update promotion in the underlying promotion repository (s.PromotionRepo)
	updatePromo, err := s.PromotionRepo.UpdatePromotionbyPromotionID(promo)

	if err != nil {
		// If no promotion is found with the given PromotionID during the update operation,
		// return a custom error indicating that a duplicate promotion was found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Promotion{}, &exception.PromotionIDNotFoundError{
				Message:     "Duplicate Promotion Found",
				PromotionID: promo.PromotionID,
			}
		}
		// If any other error occurs during the update operation, return the encountered error
		return models.Promotion{}, err
	}
	// Return the updated promotion along with a nil error if the operation is successful
	return updatePromo, nil
}

// DeletePromotionByPromotionID will delete data based on promotionID request
func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error {
	return s.PromotionRepo.DeletePromotionbyPromotionID(promotionID)
}
