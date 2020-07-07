package usecases

import (
	"hotelAPI/main/masters/models"
	"hotelAPI/main/masters/repositories"
)

//RoomUsecaseImpl app
type RoomUsecaseImpl struct {
	roomRepo repositories.RoomRepository
}

//GetAvailableRoom app
func (s RoomUsecaseImpl) GetAvailableRoom() ([]*models.Rooms, error) {
	rooms, err := s.roomRepo.SelectAvailableRoom()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

//GetBookedRoom app
func (s RoomUsecaseImpl) GetBookedRoom() ([]*models.Rooms, error) {
	rooms, err := s.roomRepo.SelectBookedRoom()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

//GetAllRoom app
func (s RoomUsecaseImpl) GetAllRoom() ([]*models.Rooms, error) {
	rooms, err := s.roomRepo.SelectAllRoom()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

//InitRoomUsecaseImpl app
func InitRoomUsecaseImpl(roomRepo repositories.RoomRepository) RoomUsecase {
	return &RoomUsecaseImpl{roomRepo}
}
