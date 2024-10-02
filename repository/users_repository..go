package repository

import "golang-crud-gin/model"

type UsersRepository interface {
	FindByemail(email string) (users model.Users, err error)
	Save(users model.Users)
}
