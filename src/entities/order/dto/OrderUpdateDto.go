package order

import (
	orderdetailsdto "skyshi.com/src/entities/order_details/dto"
)

type OrderUpdateDto struct {
	Id                 int                                    `json:"id"`
	CustomerId         int                                    `json:"customer_id" validate:"required"`
	Status             string                                 `json:"status"`
	OrderDetailsUpdate []orderdetailsdto.OrderDetailUpdateDto `json:"order_details" validate:"required"`
}
