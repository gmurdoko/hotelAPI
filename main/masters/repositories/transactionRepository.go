package repositories

import "hotelAPI/main/masters/models"

// ReserveRepository app
type ReserveRepository interface {
	// SelectAvailableReserve() ([]*models.Reserves, error)
	// SelectBookedReserve() ([]*models.Reserves, error)
	SelectAllReserve() ([]*models.Reserves, error)
	AddReserve(inReserve *models.Reserves) error
	// EditReserve(inReserve *models.Reserves) error
	// DelReserve(id int) error
}
