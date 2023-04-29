package order

import (
	"fmt"
	// "io/ioutil"
	"strconv"

	dto "skyshi.com/src/entities/order/dto"
	orderdetails "skyshi.com/src/entities/order_details"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Orderservice interface {
	GetAll() []OrderModel
	GetByCustomerId(id int) []OrderModel
	GetByID(id int) OrderModel
	Create(ctx *gin.Context) (*OrderModel, error)
	Update(ctx *gin.Context) (*OrderModel, error)
	UpdateStatusByCustomerId(id int, status string) (*[]OrderModel, error)
	Delete(ctx *gin.Context) (*OrderModel, error)
}

type OrderserviceImpl struct {
	orderRepository       OrderRepository
	orderDetailRepository orderdetails.OrderDetailRepository
}

func NewOrderservice(orderRepository OrderRepository, orderDetailRepository orderdetails.OrderDetailRepository) Orderservice {
	return &OrderserviceImpl{orderRepository, orderDetailRepository}
}

func (us *OrderserviceImpl) GetByCustomerId(id int) []OrderModel {
	return us.orderRepository.GetByCustomerId(id)
}

func (us *OrderserviceImpl) GetAll() []OrderModel {
	return us.orderRepository.FindAll()
}

func (us *OrderserviceImpl) GetByID(id int) OrderModel {
	return us.orderRepository.FindOne(id)
}

func (us *OrderserviceImpl) Create(ctx *gin.Context) (*OrderModel, error) {
	var input dto.OrderCreateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {

		fmt.Println(err.Error())
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	order := OrderModel{
		CustomerId: input.CustomerId,
		Status:     input.Status,
	}

	result, err := us.orderRepository.Save(order)

	if err != nil {
		return nil, err
	}

	var orderdetailsdata []orderdetails.OrderDetailModel

	for i := 0; i < len(input.OrderDetailsCreate); i++ {
		var orderdetail = orderdetails.OrderDetailModel{
			OrderId:   result.ID,
			ProductId: input.OrderDetailsCreate[i].ProductId,
			Qty:       input.OrderDetailsCreate[i].Qty,
			Price:     input.OrderDetailsCreate[i].Price,
			Total:     input.OrderDetailsCreate[i].Qty * input.OrderDetailsCreate[i].Price,
		}

		resultOrderDetail, err := us.orderDetailRepository.Save(orderdetail)

		if err != nil {
			return nil, err
		}

		orderdetailsdata = append(orderdetailsdata, us.orderDetailRepository.FindOne(resultOrderDetail.ID))
	}

	if err != nil {
		return nil, err
	} else {
		result.OrderDetails = orderdetailsdata

		return result, err
	}
}

func (us *OrderserviceImpl) Update(ctx *gin.Context) (*OrderModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.OrderUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	order := OrderModel{
		ID:         int(id),
		CustomerId: input.CustomerId,
		Status:     input.Status,
	}

	result, err := us.orderRepository.Update(order)

	var orderdetailsdata []orderdetails.OrderDetailModel

	_, err = us.orderDetailRepository.DeleteAll(int(id))

	if err != nil {
		return nil, err
	}

	for i := 0; i < len(input.OrderDetailsUpdate); i++ {
		var orderdetail = orderdetails.OrderDetailModel{
			OrderId:   result.ID,
			ProductId: input.OrderDetailsUpdate[i].ProductId,
			Qty:       input.OrderDetailsUpdate[i].Qty,
			Price:     input.OrderDetailsUpdate[i].Price,
			Total:     input.OrderDetailsUpdate[i].Qty * input.OrderDetailsUpdate[i].Price,
		}

		resultOrderDetail, err := us.orderDetailRepository.Save(orderdetail)

		if err != nil {
			return nil, err
		}

		orderdetailsdata = append(orderdetailsdata, us.orderDetailRepository.FindOne(resultOrderDetail.ID))
	}

	if err != nil {
		return nil, err
	} else {
		result.OrderDetails = orderdetailsdata

		return result, err
	}
}

func (us *OrderserviceImpl) UpdateStatusByCustomerId(id int, status string) (*[]OrderModel, error) {
	result, err := us.orderRepository.UpdateStatusByCustomerId(id, status)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *OrderserviceImpl) Delete(ctx *gin.Context) (*OrderModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	order := OrderModel{
		ID: int(id),
	}

	result, err := us.orderRepository.Delete(order)

	if err != nil {
		return nil, err
	}

	return result, nil
}
