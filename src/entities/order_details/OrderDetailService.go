package orderdetails

import (
	"fmt"
	// "io/ioutil"
	"strconv"

	dto "skyshi.com/src/entities/order_details/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type OrderDetailservice interface {
	GetAll() []OrderDetailModel
	GetAllPendingOrders() []OrderDetailModel
	GetByID(id int) OrderDetailModel
	Create(ctx *gin.Context) (*OrderDetailModel, error)
	Update(ctx *gin.Context) (*OrderDetailModel, error)
	Delete(ctx *gin.Context) (*OrderDetailModel, error)
}

type OrderDetailserviceImpl struct {
	orderdetailRepository OrderDetailRepository
}

func NewOrderDetailservice(orderdetailRepository OrderDetailRepository) OrderDetailservice {
	return &OrderDetailserviceImpl{orderdetailRepository}
}

func (us *OrderDetailserviceImpl) GetAll() []OrderDetailModel {
	return us.orderdetailRepository.FindAll()
}

func (us *OrderDetailserviceImpl) GetAllPendingOrders() []OrderDetailModel {
	return us.orderdetailRepository.FindAllPendingOrders()
}

func (us *OrderDetailserviceImpl) GetByID(id int) OrderDetailModel {
	return us.orderdetailRepository.FindOne(id)
}

func (us *OrderDetailserviceImpl) Create(ctx *gin.Context) (*OrderDetailModel, error) {
	var input dto.OrderDetailCreateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {

		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(input)

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	orderdetail := OrderDetailModel{
		ProductId: input.ProductId,
		Qty:       input.Qty,
		Price:     input.Price,
		Total:     input.Total,
	}

	result, err := us.orderdetailRepository.Save(orderdetail)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *OrderDetailserviceImpl) Update(ctx *gin.Context) (*OrderDetailModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.OrderDetailUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	orderdetail := OrderDetailModel{
		ID:        int(id),
		ProductId: input.ProductId,
		Qty:       input.Qty,
		Price:     input.Price,
		Total:     input.Total,
	}

	result, err := us.orderdetailRepository.Update(orderdetail)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *OrderDetailserviceImpl) Delete(ctx *gin.Context) (*OrderDetailModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	orderdetail := OrderDetailModel{
		ID: int(id),
	}

	result, err := us.orderdetailRepository.Delete(orderdetail)

	if err != nil {
		return nil, err
	}

	return result, nil
}
