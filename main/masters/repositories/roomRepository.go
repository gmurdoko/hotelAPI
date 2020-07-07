package repositories

import "hotelAPI/main/masters/models"

// RoomRepository app
type RoomRepository interface {
	SelectAvailableRoom() ([]*models.Rooms, error)
	SelectBookedRoom() ([]*models.Rooms, error)
	SelectAllRoom() ([]*models.Rooms, error)
	AddRoom(inRoom *models.Rooms) error
	EditRoom(inRoom *models.Rooms) error
	DelRoom(id int) error
}
