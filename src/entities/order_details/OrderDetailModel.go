package orderdetails

import (
	"gorm.io/gorm"

	product "skyshi.com/src/entities/product"
)

type OrderDetailModel struct {
	ID        int `gorm:"AUTO_INCREMENT"`
	OrderId   int `json:"order_id" gorm:"type:int(10)"`
	ProductId int `json:"product_id" gorm:"type:int(10)"`
	Qty       int `json:"qty" gorm:"type:int(10)"`
	Price     int `json:"price" gorm:"type:int(10)"`
	Total     int `json:"total" gorm:"type:int(10)"`
	
	gorm.Model
	Product product.ProductModel `gorm:"foreignKey:ProductId"`
}

func (OrderDetailModel) TableName() string {
	return "order_details"
}
