package controller

import (
	dto "crud-gin/dtos"
	model "crud-gin/models"
	repository "crud-gin/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	Repository repository.IUsersRepository
}

type IUsersController interface {
	FindOneById(ctx *gin.Context)
	FindMany(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Update(ctx *gin.Context)
}

func CreateUsersController(Repository repository.IUsersRepository) IUsersController {
	return &UsersController{Repository: Repository}
}

// @Summary			Get user by id
// @Param			user_id path string true "asd"
// @Produce			application/json
// @Tags			users
// @Router			/users/{user_id} [get]
// @Success			200 {object} model.User
func (controller *UsersController) FindOneById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, err := controller.Repository.FindOneById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorDto{Status: 404, Error_code: "user_not_found"})
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, user)
}

// @Summary			Delete user by id
// @Param			user_id path string true "user id"
// @Produce			application/json
// @Tags			users
// @Router			/users/{user_id} [delete]
// @Success			204 ""
func (controller *UsersController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	_, err := controller.Repository.FindOneById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorDto{Status: 404, Error_code: "user_not_found"})
	}

	controller.Repository.Delete(id)
	ctx.Header("Content-Type", "application/json")
	ctx.Status(http.StatusNoContent)
}

// @Summary			Get users
// @Produce			application/json
// @Tags			users
// @Router			/users [get]
// @Success			200 {object} []model.User
func (controller *UsersController) FindMany(ctx *gin.Context) {
	users := controller.Repository.FindMany()
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, users)
}

// @Summary			Create user
// @Produce			application/json
// @Tags			users
// @Router			/users [post]
// @Param			user body dto.CreateUserDto true "create user"
// @Success			201 {object} model.User
func (controller *UsersController) Create(ctx *gin.Context) {
	body := dto.CreateUserDto{}
	err := ctx.ShouldBindJSON(&body)
	ctx.Header("Content-Type", "application/json")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user := model.User{Name: body.Name, Secondname: body.Secondname, Age: body.Age}
	createdUser := controller.Repository.Create(user)
	ctx.JSON(http.StatusCreated, createdUser)
}

// @Summary			User user by id
// @Produce			application/json
// @Tags			users
// @Router			/users/{user_id} [patch]
// @Param			user_id path string true "user id"
// @Param			user body dto.UpdateUserDto true "update user"
// @Success			200 {object} model.User
func (controller *UsersController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	_, err := controller.Repository.FindOneById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorDto{Status: 404, Error_code: "user_not_found"})
	}

	body := dto.UpdateUserDto{}
	ctx.ShouldBindJSON(&body)
	user := controller.Repository.Update(model.User{Name: body.Name, Secondname: body.Secondname, Age: body.Age, Id: id})
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, user)
}
