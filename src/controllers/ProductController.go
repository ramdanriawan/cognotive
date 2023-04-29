package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"github.com/gin-gonic/gin"
	customer "skyshi.com/src/entities/customer"
	product "skyshi.com/src/entities/product"
)

type ProductController struct {
	productservice product.Productservice
	customerservice customer.CustomerService
	ctx            *gin.Context
}

func NewProductController(productservice product.Productservice, customerservice customer.CustomerService, ctx *gin.Context) ProductController {
	return ProductController{productservice, customerservice, ctx}
}

func (uc *ProductController) Index(ctx *gin.Context) {

	type DayAndTime struct {
	}

	data := uc.productservice.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *ProductController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.productservice.GetByID(int64(id))

	type DayAndTime struct {
	}

	days := []*DayAndTime{}

	if data.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Product with ID %d Not Found", id),
			"data":    days,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *ProductController) Create(ctx *gin.Context) {

	data, err := uc.productservice.Create(ctx)

	if err != nil {

		if strings.Contains(err.Error(), "Name") {
			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": "name cannot be null",
			})

		} else if strings.Contains(err.Error(), "CustomerId") {
			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": "customer_group_id cannot be null",
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

func (uc *ProductController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	productModel := uc.productservice.GetByID(int64(id))

	if productModel.ID < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Product with ID %d Not Found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.productservice.Update(ctx)

	if err != nil {

		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    uc.productservice.GetByID(int64(id)),
	})
}

func (uc *ProductController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	productModel := uc.productservice.GetByID(int64(id))

	if productModel.ID < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Product with ID %d Not Found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.productservice.Delete(ctx)

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
