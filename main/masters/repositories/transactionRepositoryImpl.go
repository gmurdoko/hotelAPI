package repositories

import (
	"database/sql"
	"hotelAPI/main/masters/models"
	"log"
	"time"
)

//ReserveRepoImpl app
type ReserveRepoImpl struct {
	db *sql.DB
}

//SelectAllReserve app
func (s ReserveRepoImpl) SelectAllReserve() ([]*models.Reserves, error) {
	data, err := s.db.Query("select mc.nik, mc.customer_name, mr.id, mr.room_name, dr.booked_at, dr.ended_at from m_customers mc join d_rooms dr on mc.id = dr.id_customer join m_rooms mr on mr.id = dr.id_room;")
	if err != nil {
		return nil, err
	}
	defer data.Close()
	var result = []*models.Reserves{}
	for data.Next() {
		var reserve = models.Reserves{}
		var err = data.Scan(&reserve.Nik, &reserve.CustomerName, &reserve.RoomID, &reserve.RoomName, &reserve.BookedAt, &reserve.EndedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, &reserve)
	}
	if err = data.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

//AddReserve app
func (s ReserveRepoImpl) AddReserve(inReserve *models.Reserves) error {
	// inReserve.BookedAt = time.Now().Format(`2006-01-02 15:04:05`)
	// inReserve.EndedAt = time.Now().Format(`2006-01-02 15:04:05`)
	CreatedAt := time.Now().Format(`2006-01-02 15:04:05`)
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	query := "insert into m_customers(nik, customer_name, created_at) values (?,?,?);"
	res, err := tx.Exec(query, inReserve.Nik, inReserve.CustomerName, CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	//insert into prices (Reserve_id, price, created_at) values (3,300000,now());
	idCustomer := int(lastID)
	query = "insert into d_rooms(id_room, id_customer, booked_at, ended_at) values (?,?,?,?);"
	res, err = tx.Exec(query, inReserve.RoomID, idCustomer, inReserve.BookedAt, inReserve.EndedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = "update m_rooms set status = 'B', edited_at = now() where id = ?;"
	_, err = tx.Exec(query, inReserve.RoomID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

//InitReserveRepoImpl app
func InitReserveRepoImpl(db *sql.DB) ReserveRepository {
	return &ReserveRepoImpl{db}
}
