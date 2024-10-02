package controller

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/service"
	"golang-crud-gin/utils"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService service.UsersService
}

func NewUsersController(service service.UsersService) *UsersController {
	return &UsersController{
		usersService: service,
	}
}

func (controller *UsersController) Create(ctx *gin.Context) {
	log.Info().Msg("create users")
	CreateUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&CreateUsersRequest)

	helper.ErrorPanic(err)

	controller.usersService.Create(CreateUsersRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// func (controller *UsersController) FindByEmail(ctx *gin.Context) {
// 	log.Info().Msg("findbyemail tags")
// 	tagId := ctx.Param("tagId")
// 	id, err := strconv.Atoi(tagId)
// 	helper.ErrorPanic(err)

// 	tagResponse := controller.tagsService.FindById(id)

// 	webResponse := response.Response{
// 		Code:   http.StatusOK,
// 		Status: "Ok",
// 		Data:   tagResponse,
// 	}
// 	ctx.Header("Content-Type", "application/json")
// 	ctx.JSON(http.StatusOK, webResponse)
// }

func (controller *UsersController) Login(ctx *gin.Context) {
	log.Info().Msg("login")
	loginUsersRequest := request.CreateUsersRequest{}
	err := ctx.ShouldBindJSON(&loginUsersRequest)
	helper.ErrorPanic(err)

	loginResponse := controller.usersService.FindByemail(loginUsersRequest.Email)
	isPasswordvalid := utils.CheckPasswordHash(loginUsersRequest.Password, loginResponse.Password)
	ctx.Header("Content-Type", "application/json")
	if !isPasswordvalid {
		webResponse := response.Response{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
			Data:   "Invalid email or password",
		}
		ctx.JSON(http.StatusUnauthorized, webResponse)
		return
	}

	token, err := utils.GenerateToken(loginResponse.Email, int64(loginResponse.Id))
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   gin.H{"message": "Login successful!", "token": token},
	}
	ctx.JSON(http.StatusOK, webResponse)
	// context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": "token"})
}
