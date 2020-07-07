package validation

import (
	"errors"
	"hotelAPI/config"
)

//ValidateRoomExist app
func ValidateRoomExist(id int) error {
	db := config.EnvConn()
	// var category = new(models.Categories)
	var status string
	err := db.QueryRow("select status from m_rooms where id =1").Scan(&status)
	if err != nil {
		return errors.New("Room didn't exist")
	}
	return nil

}

//ValidateRoomAvailable app
func ValidateRoomAvailable(id int) error {
	db := config.EnvConn()
	// var category = new(models.Categories)
	var status string
	err := db.QueryRow("select status from m_rooms where id =?", id).Scan(&status)
	if err != nil {
		return errors.New("Room didn't exist")
	}
	if status == "B" {
		return errors.New("Room already booked")
	} else if status == "D" {
		return errors.New("Room was deleted")
	}
	return nil

}
