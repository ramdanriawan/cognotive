package orderdetails

import (
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	FindAll() []OrderDetailModel
	FindAllPendingOrders() []OrderDetailModel
	FindOne(id int) OrderDetailModel
	Save(orderdetail OrderDetailModel) (*OrderDetailModel, error)
	Update(orderdetail OrderDetailModel) (*OrderDetailModel, error)
	Delete(orderdetail OrderDetailModel) (*OrderDetailModel, error)
	DeleteAll(orderId int) (*bool, error)
}

type OrderDetailRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &OrderDetailRepositoryImpl{db}
}

func (ur *OrderDetailRepositoryImpl) FindAll() []OrderDetailModel {
	var orderdetails []OrderDetailModel

	_ = ur.db.Preload("Product").Find(&orderdetails)

	return orderdetails

}

func (ur *OrderDetailRepositoryImpl) FindAllPendingOrders() []OrderDetailModel {
	var orderdetails []OrderDetailModel

	_ = ur.db.Preload("Product").Where("status", "Pending").Find(&orderdetails)

	return orderdetails

}

func (ur *OrderDetailRepositoryImpl) FindOne(id int) OrderDetailModel {
	var orderdetail OrderDetailModel
	_ = ur.db.Preload("Product").Find(&orderdetail, id)

	return orderdetail
}

func (ur *OrderDetailRepositoryImpl) Save(orderdetail OrderDetailModel) (*OrderDetailModel, error) {
	result := ur.db.Save(&orderdetail)

	if result.Error != nil {
		return nil, result.Error
	}

	return &orderdetail, nil
}

func (ur *OrderDetailRepositoryImpl) Update(orderdetail OrderDetailModel) (*OrderDetailModel, error) {

	result := ur.db.Model(&orderdetail).Updates(&orderdetail)

	if result.Error != nil {

		return nil, result.Error
	}

	return &orderdetail, nil
}

func (ur *OrderDetailRepositoryImpl) Delete(orderdetail OrderDetailModel) (*OrderDetailModel, error) {
	result := ur.db.Preload("Product").Delete(&orderdetail)

	if result.Error != nil {
		return nil, result.Error
	}

	return &orderdetail, nil
}

func (ur *OrderDetailRepositoryImpl) DeleteAll(orderId int) (*bool, error) {
	result := ur.db.Preload("Product").Delete(&OrderDetailModel{}, &OrderDetailModel{OrderId: orderId});

	isFalse := false
	isTrue := true

	if result.Error != nil {
		return &isFalse, result.Error
	}

	return &isTrue, nil
}
