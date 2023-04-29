package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	customer "skyshi.com/src/entities/customer"
	order "skyshi.com/src/entities/order"
)

type CustomerController struct {
	customerService     customer.CustomerService
	orderservice        order.Orderservice
	orderRepository     order.OrderRepository
	orderRepositoryImpl order.OrderRepositoryImpl
	ctx                 *gin.Context
}

func NewCustomerController(customerService customer.CustomerService, orderservice order.Orderservice, orderRepository order.OrderRepository, orderRepositoryImpl order.OrderRepositoryImpl, ctx *gin.Context) CustomerController {
	return CustomerController{customerService, orderservice, orderRepository, orderRepositoryImpl, ctx}
}

func (uc *CustomerController) Index(ctx *gin.Context) {

	data := uc.customerService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *CustomerController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.customerService.GetByID(id)

	fmt.Println(id)
	fmt.Println("53636345635346346334524525")

	if data.ID == 0 {
		ctx.JSON(404, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Customer with ID %d Not Found", id),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *CustomerController) Create(ctx *gin.Context) {
	data, err := uc.customerService.Create(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Name") {

			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": "name cannot be null",
			})
		} else {
			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": err.Error(),
			})
		}

		ctx.Abort()

		return
	}

	ctx.JSON(201, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *CustomerController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	customerModel := uc.customerService.GetByID(id)

	if customerModel.ID < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Customer with ID %d Not Found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.customerService.Update(ctx)

	if err != nil {

		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})

		ctx.Abort()

		return
	}

	ctx.JSON(200, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    uc.customerService.GetByID(id),
	})
}

func (uc *CustomerController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	customerModel := uc.customerService.GetByID(id)
	fmt.Println(id)
	fmt.Println("idsdfojdsfjdlsfjdklsfjkdsjfdiklshjfklidshfkldshfjkdhsfjkhsdfjkhsdjkfghdjksfgdjksgfhjdsf")
	if customerModel.ID < 1 {
		ctx.JSON(404, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Customer with ID %d Not Found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.customerService.Delete(ctx)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})

		ctx.Abort()

		return

	}

	type ResponseData struct {
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    ResponseData{},
	})
}
