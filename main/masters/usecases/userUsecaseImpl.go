package usecases

import (
	"hotelAPI/main/masters/models"
	"hotelAPI/main/masters/repositories"
	"hotelAPI/utils/pwd"
)

//UserUsecaseImpl app
type UserUsecaseImpl struct {
	UserRepo repositories.UserRepository
}

//GetAuthUsers app
func (s UserUsecaseImpl) GetAuthUsers(inUser *models.Users) (bool, error) {
	users, err := s.UserRepo.GetAuth(inUser)
	if err != nil {
		return false, err
	}
	isPwdValid := pwd.CheckPasswordHash(inUser.Password, users.Password)
	if inUser.UserName != users.UserName || !isPwdValid {
		return false, err
	}
	return true, nil
}

//PostUser app
func (s UserUsecaseImpl) PostUser(inUser *models.Users) error {
	err := s.UserRepo.AddUser(inUser)
	if err != nil {
		return err
	}
	return nil
}

//InitUserUsecaseImpl app
func InitUserUsecaseImpl(userRepo repositories.UserRepository) UserUsecase {
	return &UserUsecaseImpl{userRepo}
}
