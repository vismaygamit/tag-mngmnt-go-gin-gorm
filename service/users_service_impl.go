package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/repository"
	"golang-crud-gin/utils"

	"github.com/go-playground/validator/v10"
)

type UsersServiceImpl struct {
	UsersRepository repository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersServiceImpl(userRepository repository.UsersRepository, validate *validator.Validate) UsersService {
	return &UsersServiceImpl{
		UsersRepository: userRepository,
		Validate:        validate,
	}
}

// Create implements TagsService
func (u *UsersServiceImpl) Create(users request.CreateUsersRequest) {
	err := u.Validate.Struct(users)
	helper.ErrorPanic(err)
	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)
	userModel := model.Users{
		Email:    users.Email,
		Password: hashedPassword,
	}
	u.UsersRepository.Save(userModel)
}

// Delete implements TagsService
// func (t *TagsServiceImpl) Delete(tagsId int) {
// 	t.TagsRepository.Delete(tagsId)
// }

// FindAll implements TagsService
// func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
// 	result := t.TagsRepository.FindAll()

// 	var tags []response.TagsResponse
// 	for _, value := range result {
// 		tag := response.TagsResponse{
// 			Id:   value.Id,
// 			Name: value.Name,
// 		}
// 		tags = append(tags, tag)
// 	}

// 	return tags
// }

// FindById implements TagsService
func (u *UsersServiceImpl) FindByemail(email string) response.UsersResponse {
	userData, err := u.UsersRepository.FindByemail(email)
	helper.ErrorPanic(err)
	userResponse := response.UsersResponse{
		// Id:    userData.Id,
		Email:    userData.Email,
		Password: userData.Password,
	}
	return userResponse
}

// Update implements TagsService
// func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
// 	tagData, err := t.TagsRepository.FindById(tags.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = tags.Name
// 	t.TagsRepository.Update(tagData)
// }
