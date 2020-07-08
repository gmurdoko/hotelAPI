package repositories

import (
	"database/sql"
	"hotelAPI/main/masters/models"
	"hotelAPI/utils/pwd"
)

//UserRepoImpl app
type UserRepoImpl struct {
	db *sql.DB
}

//GetAuth app
func (u *UserRepoImpl) GetAuth(userIn *models.Users) (*models.Users, error) {

	row := u.db.QueryRow(`select username,password from m_users where username=?`, userIn.UserName)
	var user = new(models.Users)
	err := row.Scan(&user.UserName, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil

}

//AddUser app
func (u *UserRepoImpl) AddUser(userIn *models.Users) error {
	hashPwd, err := pwd.HashPassword(userIn.Password)
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec(`insert into m_users (username, password) values (?,?)`, userIn.UserName, hashPwd)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	userIn.Password = hashPwd
	return nil

}

//InitUserRepoImpl app
func InitUserRepoImpl(db *sql.DB) UserRepository {
	return &UserRepoImpl{db}
}
