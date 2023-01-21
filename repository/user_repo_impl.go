package repository

import (
	"fmt"
	"gorm.io/gorm"
	"medastra/models"
)

const UserTable = "user"
const LoginTable = "login"
const slotTable = "slots"
const doctorTable = "doctor"

type RepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		Db: Db,
	}
}
func (u RepositoryImpl) SaveProfile(user *models.UserProfile) error {
	err := u.Db.Table(UserTable).Create(&user).Error
	return err
}
func (u RepositoryImpl) UpdateProfile(user *models.UserProfile) error {
	err := u.Db.Table(UserTable).Model(&user).Updates(&user).Error
	return err
}

func (u RepositoryImpl) DeleteProfile(Id int) error {
	var user models.UserProfile
	err := u.Db.Table(UserTable).Where("id=?", Id).Delete(&user).Error
	return err
}
func (u RepositoryImpl) FindProfileByUserName(username string) (*models.UserProfile, error) {
	var result models.UserProfile
	err := u.Db.Table(UserTable).
		Where("username = ?", username).
		First(&result).
		Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}
func (u RepositoryImpl) CreateLogin(username string, password string) error {
	var userLogin models.User
	err := u.Db.Table(LoginTable).Create(&userLogin).Error
	return err
}
func (u RepositoryImpl) CheckLogin(username string, password string) error {
	var userLogin models.User
	err := u.Db.Table(LoginTable).Where("username=?", username).First(&userLogin).Error
	if err != nil {
		return err
	}
	if userLogin.Password != password {
		return fmt.Errorf("Incorrect password or incorrect username")
	}
	return nil
}
