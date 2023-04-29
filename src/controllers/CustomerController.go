package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	golangjwt "github.com/golang-jwt/jwt"
	customer "skyshi.com/src/entities/customer"
	customerdto "skyshi.com/src/entities/customer/dto"
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

func (uc *CustomerController) Authenticate(ctx *gin.Context) {
	var sampleSecretKey = []byte("SecretYouShouldHide")

	token := golangjwt.New(golangjwt.SigningMethodHS256)

	claims := token.Claims.(golangjwt.MapClaims)

	var input customerdto.CustomerAuthenticateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})

		return
	}

	var customer = uc.customerService.GetByEmailAndPassword(input.Email, input.Password)

	if customer.ID < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Error",
			"message": "Customer not found",
		})

		return
	}

	claims["id"] = customer.ID
	claims["exp"] = time.Now().Second() * 3600 * 12

	tokenString, err := token.SignedString(sampleSecretKey)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    tokenString,
	})
}

func (uc *CustomerController) decodeUserIdByToken(user_token string) int {
	parsedToken, _ := jwt.Parse(user_token, nil)

	claims, _ := parsedToken.Claims.(jwt.MapClaims)
	
	if claims["id"] == nil {
		return -1
	}

	// var theExpired, err = time.Parse("2006-Jan-02", claims["exp"]);
	
	// if err != nil {
	// 	return -1
	// }

	

	id := claims["id"].(float64)

	return int(id)
}

func (uc *CustomerController) Index(ctx *gin.Context) {
	admin_token := ctx.Query("admin_token")

	user_id := int(uc.decodeUserIdByToken(admin_token))

	if user_id != 0 {

		type DayAndTime struct {
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Error",
			"message": "Admin token not found!",
			"data":    DayAndTime{},
		})

		return
	}

	data := uc.customerService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *CustomerController) GetByID(ctx *gin.Context) {
	admin_token := ctx.Query("admin_token")

	user_id := int(uc.decodeUserIdByToken(admin_token))

	if user_id != 0 {

		type DayAndTime struct {
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Error",
			"message": "Admin token not found!",
			"data":    DayAndTime{},
		})

		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.customerService.GetByID(id)

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
	admin_token := ctx.Query("admin_token")

	user_id := int(uc.decodeUserIdByToken(admin_token))

	if user_id != 0 {

		type DayAndTime struct {
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Error",
			"message": "Admin token not found!",
			"data":    DayAndTime{},
		})

		return
	}

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
	admin_token := ctx.Query("admin_token")

	user_id := int(uc.decodeUserIdByToken(admin_token))

	if user_id != 0 {

		type DayAndTime struct {
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Error",
			"message": "Admin token not found!",
			"data":    DayAndTime{},
		})

		return
	}

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
	admin_token := ctx.Query("admin_token")

	user_id := int(uc.decodeUserIdByToken(admin_token))

	if user_id != 0 {

		type DayAndTime struct {
		}

		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Error",
			"message": "Admin token not found!",
			"data":    DayAndTime{},
		})

		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	customerModel := uc.customerService.GetByID(id)

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
