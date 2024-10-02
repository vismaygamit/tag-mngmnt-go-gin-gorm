package repository

import (
	"errors"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersREpositoryImpl(Db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{Db: Db}
}

func (u *UsersRepositoryImpl) Save(users model.Users) {
	result := u.Db.Create(&users)
	helper.ErrorPanic(result.Error)
}

func (u *UsersRepositoryImpl) FindByemail(email string) (users model.Users, err error) {
	var user model.Users
	result := u.Db.Where("email = ?", email).First(&user)
	if result != nil {
		return user, nil
	} else {
		return user, errors.New("Invalid email or password")
	}
}
