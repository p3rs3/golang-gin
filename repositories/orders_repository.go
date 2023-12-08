package repository

import (
	model "crud-gin/models"

	"gorm.io/gorm"
)

type IOrdersRepository interface {
	FindOneById(id int) (model.Order, error)
	FindMany() []model.Order
	Create(order model.Order) model.Order
	Delete(id int) error
	Update(order model.Order) model.Order
}

type OrdersRepository struct {
	Db *gorm.DB
}

func CreateOrdersRepository(Db *gorm.DB) IOrdersRepository {
	return &OrdersRepository{Db: Db}
}

func (or *OrdersRepository) FindOneById(id int) (order model.Order, err error) {
	result := or.Db.First(&order, id)

	if result.Error != nil {
		err = result.Error
	}

	return
}

func (or *OrdersRepository) FindMany() (orders []model.Order) {
	or.Db.Limit(USERS_LIMIT).Find(&orders)
	return orders
}

func (or *OrdersRepository) Create(order model.Order) model.Order {
	or.Db.Create(&order)
	return order
}

func (or *OrdersRepository) Update(order model.Order) model.Order {
	or.Db.Model(&order).Updates(order)
	return order
}

func (or *OrdersRepository) Delete(id int) (err error) {
	result := or.Db.Delete(&model.Order{}, id)

	if result != nil {
		err = result.Error
	}

	return
}
