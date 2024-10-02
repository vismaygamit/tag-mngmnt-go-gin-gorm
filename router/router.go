package router

import (
	"golang-crud-gin/controller"
	"golang-crud-gin/middlewares"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(tagsController *controller.TagsController, usersController *controller.UsersController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	baseRouter.GET("/tags", tagsController.FindAll)
	baseRouter.GET("/tags/:tagId", tagsController.FindById)
	tagsRouter.Use(middlewares.Authenticate)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	baseRouter.POST("/signup", usersController.Create)
	baseRouter.POST("/login", usersController.Login)
	return router
}
