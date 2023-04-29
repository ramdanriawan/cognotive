package order

import (
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll() []OrderModel
	GetByCustomerId(id int) []OrderModel
	FindOne(id int) OrderModel
	Save(order OrderModel) (*OrderModel, error)
	Update(order OrderModel) (*OrderModel, error)
	UpdateStatusByCustomerId(id int, status string) (*[]OrderModel, error)
	Delete(order OrderModel) (*OrderModel, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db}
}

func (ur *OrderRepositoryImpl) GetByCustomerId(id int) []OrderModel {
	var orders []OrderModel

	_ = ur.db.Preload("OrderDetails.Product").Where("customer_id", id).Find(&orders)

	return orders
}

func (ur *OrderRepositoryImpl) FindAll() []OrderModel {
	var orders []OrderModel

	_ = ur.db.Preload("OrderDetails.Product").Find(&orders)

	return orders

}

func (ur *OrderRepositoryImpl) FindOne(id int) OrderModel {
	var order OrderModel
	_ = ur.db.Preload("OrderDetails.Product").Find(&order, id)

	return order
}

func (ur *OrderRepositoryImpl) Save(order OrderModel) (*OrderModel, error) {
	result := ur.db.Save(&order)

	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (ur *OrderRepositoryImpl) Update(order OrderModel) (*OrderModel, error) {

	result := ur.db.Model(&order).Updates(&order)

	if result.Error != nil {

		return nil, result.Error
	}

	return &order, nil
}

func (ur *OrderRepositoryImpl) UpdateStatusByCustomerId(id int, status string) (*[]OrderModel, error) {

	result := ur.db.Model(&OrderModel{CustomerId: id}).Where("customer_id", id).Updates(OrderModel{Status: status})

	if result.Error != nil {

		return nil, result.Error
	}

	data := ur.GetByCustomerId(id)

	if result.Error != nil {

		return nil, result.Error
	}

	return &data, nil
}

func (ur *OrderRepositoryImpl) Delete(order OrderModel) (*OrderModel, error) {
	result := ur.db.Delete(&order)

	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}
