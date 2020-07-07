package repositories

import (
	"database/sql"
	"hotelAPI/main/masters/models"
	"log"
)

//RoomRepoImpl app
type RoomRepoImpl struct {
	db *sql.DB
}

//SelectAvailableRoom app
func (s RoomRepoImpl) SelectAvailableRoom() ([]*models.Rooms, error) {
	data, err := s.db.Query("select mr.id, mr.room_name, p.price, mr.status, mr.created_at, mr.edited_at from m_rooms mr join prices p on mr.id = p.room_id where mr.status ='A';")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var result = []*models.Rooms{}
	for data.Next() {
		var room = models.Rooms{}
		var err = data.Scan(&room.ID, &room.RoomName, &room.Price, &room.Status, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, &room)
	}
	if err = data.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

//SelectAllRoom app
func (s RoomRepoImpl) SelectAllRoom() ([]*models.Rooms, error) {
	data, err := s.db.Query("select mr.id, mr.room_name, p.price, mr.status, mr.created_at, mr.edited_at from m_rooms mr join prices p on mr.id = p.room_id;")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var result = []*models.Rooms{}
	for data.Next() {
		var room = models.Rooms{}
		var err = data.Scan(&room.ID, &room.RoomName, &room.Price, &room.Status, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, &room)
	}
	if err = data.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

//SelectBookedRoom app
func (s RoomRepoImpl) SelectBookedRoom() ([]*models.Rooms, error) {
	data, err := s.db.Query("select mr.id, mr.room_name, p.price, mr.status, mr.created_at, mr.edited_at from m_rooms mr join prices p on mr.id = p.room_id where mr.status ='B';")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var result = []*models.Rooms{}
	for data.Next() {
		var room = models.Rooms{}
		var err = data.Scan(&room.ID, &room.RoomName, &room.Price, &room.Status, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, &room)
	}
	if err = data.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

//InitRoomRepoImpl app
func InitRoomRepoImpl(db *sql.DB) RoomRepository {
	return &RoomRepoImpl{db}
}
