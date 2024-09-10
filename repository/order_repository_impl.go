package repository

import (
	"hacktiv8_golang_assignment_2/entity"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) CreateOrder(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepositoryImpl) GetOrders() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *OrderRepositoryImpl) GetOrderById(id uint) (*entity.Order, error) {
	var order entity.Order
	err := r.db.Preload("Items").First(&order, id).Error
	return &order, err
}

func (r *OrderRepositoryImpl) UpdateOrder(order *entity.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderRepositoryImpl) DeleteOrder(id uint) error {
	return r.db.Delete(&entity.Order{}, id).Error
}
