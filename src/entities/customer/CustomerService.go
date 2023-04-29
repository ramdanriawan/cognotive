package customer

import (
	"strconv"

	dto "skyshi.com/src/entities/customer/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerService interface {
	GetAll() []CustomerModel
	GetByID(id int) CustomerModel
	Create(ctx *gin.Context) (*CustomerModel, error)
	Update(ctx *gin.Context) (*CustomerModel, error)
	Delete(ctx *gin.Context) (*CustomerModel, error)
}

type CustomerServiceImpl struct {
	customerRepository CustomerRepository
}

func NewCustomerService(ordertransactionRepository CustomerRepository) CustomerService {
	return &CustomerServiceImpl{ordertransactionRepository}
}

func (us *CustomerServiceImpl) GetAll() []CustomerModel {
	return us.customerRepository.FindAll()
}

func (us *CustomerServiceImpl) GetByID(id int) CustomerModel {
	return us.customerRepository.FindOne(id)
}

func (us *CustomerServiceImpl) Create(ctx *gin.Context) (*CustomerModel, error) {
	var input dto.CustomerCreateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	customer := CustomerModel{
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}

	result, err := us.customerRepository.Save(customer)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *CustomerServiceImpl) Update(ctx *gin.Context) (*CustomerModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.CustomerUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	customer := CustomerModel{
		ID:       int(id),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	if input.Password == "" {
		customer = CustomerModel{
			ID:    int(id),
			Name:  input.Name,
			Email: input.Email,
		}
	}

	result, err := us.customerRepository.Update(customer)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *CustomerServiceImpl) Delete(ctx *gin.Context) (*CustomerModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	ordertransaction := CustomerModel{
		ID: int(id),
	}

	result, err := us.customerRepository.Delete(ordertransaction)

	if err != nil {
		return nil, err
	}

	return result, nil
}
