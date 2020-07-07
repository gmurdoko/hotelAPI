package usecases

import (
	"hotelAPI/main/masters/models"
)

//RoomUsecase app
type RoomUsecase interface {
	GetAvailableRoom() ([]*models.Rooms, error)
	GetBookedRoom() ([]*models.Rooms, error)
	GetAllRoom() ([]*models.Rooms, error)
	// GetRoom(id int) (*models.Categories, error)
	// DeleteRoom(id int) error
	// PostRoom(inRoom *models.Categories) error
	// PutRoom(inRoom *models.Categories) error
}
