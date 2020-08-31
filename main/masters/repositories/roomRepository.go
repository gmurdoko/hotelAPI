package repositories

import "hotelAPI/main/masters/models"

// RoomRepository app
type RoomRepository interface {
	SelectAvailableRoom() ([]*models.Rooms, error)
	SelectBookedRoom() ([]*models.Rooms, error)
	SelectAllRoom(keyword, offset, limit, status, orderBy, sort string) ([]*models.Rooms, *int, error)
	AddRoom(inRoom *models.Rooms) error
	EditRoom(inRoom *models.Rooms) error
	DelRoom(id int) error
	SelectRoom(id int) (*models.Rooms, error)
}
