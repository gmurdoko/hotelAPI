package repositories

import (
	"database/sql"
	"fmt"
	"hotelAPI/main/masters/models"
	"log"
	"time"
)

//RoomRepoImpl app
type RoomRepoImpl struct {
	db *sql.DB
}

//SelectAvailableRoom app
func (s RoomRepoImpl) SelectAvailableRoom() ([]*models.Rooms, error) {
	data, err := s.db.Query("select mr.id, mr.room_name, p.price, mr.status, mr.created_at, mr.edited_at from m_rooms mr join prices p on mr.id = p.room_id where mr.status ='A' and p.status ='A' order by mr.id;")
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
func (s RoomRepoImpl) SelectAllRoom(keyword, offset, limit, status, orderBy, sort string) ([]*models.Rooms, *int, error) {
	queryIn := fmt.Sprintf(`select mr.id, mr.room_name, p.price, mr.status, mr.created_at, mr.edited_at from m_rooms mr join prices p on mr.id = p.room_id where p.status ='A' and mr.status= ? and (mr.room_name LIKE ? or price LIKE ?) order by %s %s limit %s , %s;`, orderBy, sort, offset, limit)
	// fmt.Println(query)
	data, err := s.db.Query(queryIn, status, "%"+keyword+"%", "%"+keyword+"%")
	// fmt.Println(s.db.Query(queryIn, status, "%"+keyword+"%", "%"+keyword+"%"))
	if err != nil {
		// fmt.Println(err, "error")
		return nil, nil, err

	}
	// fmt.Println("ini saya")

	defer data.Close()
	var result = []*models.Rooms{}
	for data.Next() {
		// fmt.Println("ini data:", data)
		var room = models.Rooms{}
		var err = data.Scan(&room.ID, &room.RoomName, &room.Price, &room.Status, &room.CreatedAt, &room.UpdatedAt)
		if err != nil {
			log.Println(err)
			return nil, nil, err
		}
		result = append(result, &room)
	}
	if err = data.Err(); err != nil {
		return nil, nil, err
	}
	var totalField = new(int)
	err = s.db.QueryRow(`select COUNT(*) from m_rooms mr join prices p on mr.id = p.room_id where p.status ='A' and mr.status=? ;`, status).Scan(&totalField)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("total Field", *totalField)
	return result, totalField, nil
}

//SelectBookedRoom app
func (s RoomRepoImpl) SelectBookedRoom() ([]*models.Rooms, error) {
	data, err := s.db.Query("select mr.id, mr.room_name, p.price, mr.status, mr.created_at, mr.edited_at from m_rooms mr join prices p on mr.id = p.room_id where mr.status ='B' and p.status ='A';")
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

//AddRoom app
func (s RoomRepoImpl) AddRoom(inRoom *models.Rooms) error {
	inRoom.CreatedAt = time.Now().Format(`2006-01-02 15:04:05`)
	inRoom.UpdatedAt = time.Now().Format(`2006-01-02 15:04:05`)
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "insert into m_rooms(room_name, created_at, edited_at) values (?,?,?);"
	res, err := tx.Exec(query, inRoom.RoomName, inRoom.CreatedAt, inRoom.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	//insert into prices (room_id, price, created_at) values (3,300000,now());
	inRoom.ID = int(lastID)
	query = "insert into prices (room_id, price, created_at, edited_at) values (?,?,?,?)"
	res, err = tx.Exec(query, inRoom.ID, inRoom.Price, inRoom.CreatedAt, inRoom.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

//EditRoom app
func (s RoomRepoImpl) EditRoom(inRoom *models.Rooms) error {
	tx, err := s.db.Begin()
	inRoom.CreatedAt = time.Now().Format(`2006-01-02 15:04:05`)
	inRoom.UpdatedAt = time.Now().Format(`2006-01-02 15:04:05`)
	if err != nil {
		return err
	}
	query := "update m_rooms set room_name = ?, status = ?, edited_at = ? where id = ?;"
	_, err = tx.Exec(query, inRoom.RoomName, inRoom.Status, inRoom.UpdatedAt, inRoom.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = "update prices set status='D', edited_at = ? where room_id = ? and status = 'A';"
	_, err = tx.Exec(query, inRoom.UpdatedAt, inRoom.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = "insert into prices (room_id, price, created_at, edited_at) values (?,?,?,?)"
	_, err = tx.Exec(query, inRoom.ID, inRoom.Price, inRoom.CreatedAt, inRoom.UpdatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//DelRoom app
func (s RoomRepoImpl) DelRoom(id int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := "update m_rooms set status = 'D', edited_at = now() where id = ?;"
	_, err = tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//InitRoomRepoImpl app
func InitRoomRepoImpl(db *sql.DB) RoomRepository {
	return &RoomRepoImpl{db}
}
