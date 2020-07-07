package usecases

import (
	"hotelAPI/main/masters/models"
)

//ReserveUsecase app
type ReserveUsecase interface {
	// GetAvailableReserve() ([]*models.Reserves, error)
	// GetBookedReserve() ([]*models.Reserves, error)
	GetAllReserve() ([]*models.Reserves, error)
	PostReserve(inReserve *models.Reserves) error
	// PutReserve(inReserve *models.Reserves) error
	// DeleteReserve(id int) error
	// GetReserve(id int) (*models.Categories, error)

}
