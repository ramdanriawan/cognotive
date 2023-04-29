package customer

import (
	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll() []CustomerModel
	FindOne(id int) CustomerModel
	FindByEmailAndPassword(email string, password string) CustomerModel
	Save(customer CustomerModel) (*CustomerModel, error)
	Update(customer CustomerModel) (*CustomerModel, error)
	Delete(customer CustomerModel) (*CustomerModel, error)
}

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{db}
}

func (ur *CustomerRepositoryImpl) FindAll() []CustomerModel {
	var customers []CustomerModel

	_ = ur.db.Preload("Orders.OrderDetails.Product").Find(&customers)

	return customers

}

func (ur *CustomerRepositoryImpl) FindOne(id int) CustomerModel {
	var customer CustomerModel
	_ = ur.db.Preload("Orders.OrderDetails.Product").Find(&customer, id)

	return customer
}

func (ur *CustomerRepositoryImpl) FindByEmailAndPassword(email string, password string) CustomerModel {
	var customer CustomerModel
	_ = ur.db.Where("email", email).Where("password", password).Find(&customer)

	return customer
}

func (ur *CustomerRepositoryImpl) Save(customer CustomerModel) (*CustomerModel, error) {
	result := ur.db.Save(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (ur *CustomerRepositoryImpl) Update(customer CustomerModel) (*CustomerModel, error) {
	result := ur.db.Model(&customer).Updates(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (ur *CustomerRepositoryImpl) Delete(customer CustomerModel) (*CustomerModel, error) {
	result := ur.db.Preload("Orders.OrderDetails.Product").Delete(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}
