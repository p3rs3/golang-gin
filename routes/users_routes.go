package routes

import (
	controller "crud-gin/controllers"

	docs "crud-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type UsersRoutes struct {
	Controller controller.IUsersController
}

type IUsersRoutes interface {
	GetRoutes() *gin.Engine
}

func CreateUsersRoutes(Controller controller.IUsersController) IUsersRoutes {
	return &UsersRoutes{Controller: Controller}
}

func (userRoutes *UsersRoutes) GetRoutes() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/users")
	tagsRouter.GET("/", userRoutes.Controller.FindMany)
	tagsRouter.POST("/", userRoutes.Controller.Create)
	tagsRouter.GET("/:id", userRoutes.Controller.FindOneById)
	tagsRouter.DELETE("/:id", userRoutes.Controller.Delete)
	tagsRouter.PATCH("/:id", userRoutes.Controller.Update)

	return router
}
