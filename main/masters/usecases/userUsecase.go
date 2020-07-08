package usecases

import "hotelAPI/main/masters/models"

//UserUsecase app
type UserUsecase interface {
	GetAuthUsers(*models.Users) (bool, error)
	PostUser(*models.Users) error
}
