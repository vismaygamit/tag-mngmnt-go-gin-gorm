package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type UsersService interface {
	Create(users request.CreateUsersRequest)
	// Update(tags request.UpdateTagsRequest)
	// Delete(tagsId int)
	// FindById(tagsId int) response.TagsResponse
	FindByemail(email string) response.UsersResponse
}
