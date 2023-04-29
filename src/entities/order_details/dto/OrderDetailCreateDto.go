package orderdetails

import (
	product "skyshi.com/src/entities/product/dto"
)

type OrderDetailCreateDto struct {
	ID        int `json:"id"`
	OrderId   int `json:"order_id" validate:"required"`
	ProductId int `json:"product_id" validate:"required"`
	Qty       int `json:"qty"  validate:"required"`
	Price     int `json:"price"`
	Total     int `json:"total"`

	ProductCreate product.ProductCreateDto
}
