package order

import (
	// "time"

	"gorm.io/gorm"
	orderdetails "skyshi.com/src/entities/order_details"
)

type OrderModel struct {
	ID         int    `gorm:"AUTO_INCREMENT"`
	CustomerId int    `json:"customer_id" gorm:"type:bigint(20)"`
	Status     string `json:"status" gorm:"type:varchar(100)"`
	Date     string `json:"date" gorm:"type:varchar(100)"`
	
	gorm.Model
	OrderDetails []orderdetails.OrderDetailModel `gorm:"foreignKey:OrderId"`
}

func (OrderModel) TableName() string {
	return "orders"
}
