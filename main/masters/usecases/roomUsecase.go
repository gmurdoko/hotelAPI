package usecases

import (
	"hotelAPI/main/masters/models"
)

//RoomUsecase app
type RoomUsecase interface {
	GetAvailableRoom() ([]*models.Rooms, error)
	GetBookedRoom() ([]*models.Rooms, error)
	GetAllRoom(keyword, offset, limit, status, orderBy, sort string) ([]*models.Rooms, *int, error)
	PostRoom(inRoom *models.Rooms) error
	PutRoom(inRoom *models.Rooms) error
	DeleteRoom(id int) error
	// GetRoom(id int) (*models.Categories, error)

}
