package repository

import "hacktiv8_golang_assignment_2/entity"

type OrderRepository interface {
	CreateOrder(order *entity.Order) error
	GetOrders() ([]entity.Order, error)
	GetOrderById(id uint) (*entity.Order, error)
	UpdateOrder(order *entity.Order) error
	DeleteOrder(id uint) error
}
