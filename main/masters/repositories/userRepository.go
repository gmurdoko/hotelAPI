package repositories

import "hotelAPI/main/masters/models"

//UserRepository app
type UserRepository interface {
	GetAuth(*models.Users) (*models.Users, error)
	AddUser(*models.Users) error
}
