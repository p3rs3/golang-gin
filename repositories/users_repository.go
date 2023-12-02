package repository

import (
	model "crud-gin/models"

	"gorm.io/gorm"
)

const USERS_LIMIT = 100

type IUsersRepository interface {
	FindOneById(id int) (model.User, error)
	FindMany() []model.User
	Create(user model.User) model.User
	Delete(id int) error
	Update(user model.User) model.User
}

type UsersRepository struct {
	Db *gorm.DB
}

func CreateUsersRepository(Db *gorm.DB) IUsersRepository {
	return &UsersRepository{Db: Db}
}

func (ur *UsersRepository) FindOneById(id int) (user model.User, err error) {
	result := ur.Db.First(&user, id)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func (ur *UsersRepository) FindMany() (users []model.User) {
	ur.Db.Limit(USERS_LIMIT).Find(&users)
	return users
}

func (ur *UsersRepository) Create(user model.User) model.User {
	ur.Db.Create(&user)
	return user
}

func (ur *UsersRepository) Update(user model.User) model.User {
	ur.Db.Save(&user)
	return user
}

func (ur *UsersRepository) Delete(id int) (err error) {
	result := ur.Db.Delete(&model.User{}, id)

	if result != nil {
		err = result.Error
	}

	return
}
