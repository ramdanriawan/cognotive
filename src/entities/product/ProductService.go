package product

import (
	"fmt"
	// "io/ioutil"
	"strconv"

	dto "skyshi.com/src/entities/product/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Productservice interface {
	GetAll() []ProductModel
	GetByID(id int64) ProductModel
	Create(ctx *gin.Context) (*ProductModel, error)
	Update(ctx *gin.Context) (*ProductModel, error)
	Delete(ctx *gin.Context) (*ProductModel, error)
}

type ProductserviceImpl struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) Productservice {
	return &ProductserviceImpl{productRepository}
}

func (us *ProductserviceImpl) GetAll() []ProductModel {
	return us.productRepository.FindAll()
}

func (us *ProductserviceImpl) GetByID(id int64) ProductModel {
	return us.productRepository.FindOne(id)
}

func (us *ProductserviceImpl) Create(ctx *gin.Context) (*ProductModel, error) {
	var input dto.ProductCreateDto

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

	product := ProductModel{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		Image:       input.Image,
	}

	result, err := us.productRepository.Save(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *ProductserviceImpl) Update(ctx *gin.Context) (*ProductModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.ProductUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	product := ProductModel{
		ID:          int(id),
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		Image:       input.Image,
	}

	result, err := us.productRepository.Update(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *ProductserviceImpl) Delete(ctx *gin.Context) (*ProductModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	product := ProductModel{
		ID: int(id),
	}

	result, err := us.productRepository.Delete(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}
