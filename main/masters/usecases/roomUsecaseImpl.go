package usecases

import (
	"fmt"
	"hotelAPI/main/masters/models"
	"hotelAPI/main/masters/repositories"
	"hotelAPI/utils/validation"
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
func (s RoomUsecaseImpl) GetAllRoom(keyword, offset, limit, status, orderBy, sort string) ([]*models.Rooms, *int, error) {
	rooms, totalField, err := s.roomRepo.SelectAllRoom(keyword, offset, limit, status, orderBy, sort)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("useCase", *totalField)
	return rooms, totalField, nil
}

//PostRoom app
func (s RoomUsecaseImpl) PostRoom(inRoom *models.Rooms) error {
	err := validation.ValidateInputNotNil(inRoom.RoomName, inRoom.Price)
	if err != nil {
		return err
	}
	err = s.roomRepo.AddRoom(inRoom)
	if err != nil {
		return err
	}
	return nil
}

//PutRoom app
func (s RoomUsecaseImpl) PutRoom(inRoom *models.Rooms) error {
	err := s.roomRepo.EditRoom(inRoom)
	if err != nil {
		return err
	}
	return nil
}

//DeleteRoom app
func (s RoomUsecaseImpl) DeleteRoom(id int) error {
	err := s.roomRepo.DelRoom(id)
	if err != nil {
		return err
	}
	return nil
}

//InitRoomUsecaseImpl app
func InitRoomUsecaseImpl(roomRepo repositories.RoomRepository) RoomUsecase {
	return &RoomUsecaseImpl{roomRepo}
}
