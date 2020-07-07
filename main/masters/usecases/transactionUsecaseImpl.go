package usecases

import (
	"hotelAPI/main/masters/models"
	"hotelAPI/main/masters/repositories"
)

//ReserveUsecaseImpl app
type ReserveUsecaseImpl struct {
	ReserveRepo repositories.ReserveRepository
}

// //GetAvailableReserve app
// func (s ReserveUsecaseImpl) GetAvailableReserve() ([]*models.Reserves, error) {
// 	Reserves, err := s.ReserveRepo.SelectAvailableReserve()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Reserves, nil
// }

// //GetBookedReserve app
// func (s ReserveUsecaseImpl) GetBookedReserve() ([]*models.Reserves, error) {
// 	Reserves, err := s.ReserveRepo.SelectBookedReserve()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return Reserves, nil
// }

//GetAllReserve app
func (s ReserveUsecaseImpl) GetAllReserve() ([]*models.Reserves, error) {
	Reserves, err := s.ReserveRepo.SelectAllReserve()
	if err != nil {
		return nil, err
	}
	return Reserves, nil
}

//PostReserve app
func (s ReserveUsecaseImpl) PostReserve(inReserve *models.Reserves) error {
	err := s.ReserveRepo.AddReserve(inReserve)
	if err != nil {
		return err
	}
	return nil
}

// //PutReserve app
// func (s ReserveUsecaseImpl) PutReserve(inReserve *models.Reserves) error {
// 	err := s.ReserveRepo.EditReserve(inReserve)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //DeleteReserve app
// func (s ReserveUsecaseImpl) DeleteReserve(id int) error {
// 	err := s.ReserveRepo.DelReserve(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

//InitReserveUsecaseImpl app
func InitReserveUsecaseImpl(ReserveRepo repositories.ReserveRepository) ReserveUsecase {
	return &ReserveUsecaseImpl{ReserveRepo}
}
