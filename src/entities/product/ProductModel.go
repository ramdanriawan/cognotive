package product

import (
	// "gorm.io/gorm"
)

type ProductModel struct {
	ID          int    `gorm:"AUTO_INCREMENT"`
	Name        string `json:"name" gorm:"type:varchar(100)"`
	Price       int  `json:"price" gorm:"type:int(10)"`
	Description string `json:"description" gorm:"type:varchar(100)"`
	Image       string `json:"image" gorm:"type:varchar(100)"`
}

func (ProductModel) TableName() string {
	return "products"
}
