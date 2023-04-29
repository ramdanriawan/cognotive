package product

import (
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() []ProductModel
	FindOne(id int64) ProductModel
	Save(product ProductModel) (*ProductModel, error)
	Update(product ProductModel) (*ProductModel, error)
	Delete(product ProductModel) (*ProductModel, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db}
}

func (ur *ProductRepositoryImpl) FindAll() []ProductModel {
	var products []ProductModel

	_ = ur.db.Find(&products)

	return products

}

func (ur *ProductRepositoryImpl) FindOne(id int64) ProductModel {
	var product ProductModel
	_ = ur.db.Find(&product, id)

	return product
}

func (ur *ProductRepositoryImpl) Save(product ProductModel) (*ProductModel, error) {
	result := ur.db.Save(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (ur *ProductRepositoryImpl) Update(product ProductModel) (*ProductModel, error) {

	result := ur.db.Model(&product).Updates(&product)

	if result.Error != nil {

		return nil, result.Error
	}

	return &product, nil
}

func (ur *ProductRepositoryImpl) Delete(product ProductModel) (*ProductModel, error) {
	result := ur.db.Delete(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}
