package customer

import (
	
	"gorm.io/gorm"

	order "skyshi.com/src/entities/order"
)

type CustomerModel struct {
	ID       int  `gorm:"AUTO_INCREMENT"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Email    string `json:"email" gorm:"type:varchar(100)"`
	Password string `json:"password" gorm:"type:varchar(100)"`
	
	gorm.Model
	Orders []order.OrderModel `gorm:"foreignKey:CustomerId"`
}

func (CustomerModel) TableName() string {
	return "customers"
}
