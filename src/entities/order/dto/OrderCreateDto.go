package order

import (
	orderdetailsdto "skyshi.com/src/entities/order_details/dto"
)

type OrderCreateDto struct {
	Id                 int                                    `json:"id"`
	CustomerId         int                                    `json:"customer_id" validate:"required"`
	Status             string                                 `json:"status"`
	OrderDetailsCreate []orderdetailsdto.OrderDetailCreateDto `json:"order_details" validate:"required"`
}
