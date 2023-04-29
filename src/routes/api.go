package route

import (

	customer "skyshi.com/src/entities/customer"
	order "skyshi.com/src/entities/order"
	order_details "skyshi.com/src/entities/order_details"
	product "skyshi.com/src/entities/product"

	controllers "skyshi.com/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx                       *gin.Context
	orderRepositoryImpl       order.OrderRepositoryImpl
	OrderDetailRepositoryImpl order_details.OrderDetailRepositoryImpl
)

func Api(router *gin.Engine, db *gorm.DB, mw gin.HandlerFunc) {
	
	orderRepository := order.NewOrderRepository(db)
	orderDetailRepository := order_details.NewOrderDetailRepository(db)
	productRepository := product.NewProductRepository(db)
	orderservice := order.NewOrderservice(orderRepository, orderDetailRepository)
	productService := product.NewProductService(productRepository)

	customerRepository := customer.NewCustomerRepository(db)
	customerService := customer.NewCustomerService(customerRepository)
	customerController := controllers.NewCustomerController(customerService, orderservice, orderRepository, orderRepositoryImpl, ctx)

	orderController := controllers.NewOrderController(orderservice, customerService, ctx)
	productController := controllers.NewProductController(productService, customerService, ctx)

	v1 := router.Group("/")
	{

		v1.POST("customer/authenticate", mw, customerController.Authenticate)
		v1.GET("customer", mw, customerController.Index)
		v1.POST("customer", mw, customerController.Create)
		v1.GET("customer/:id", mw, customerController.GetByID)
		v1.PATCH("customer/:id", mw, customerController.Update)
		v1.DELETE("customer/:id", mw, customerController.Delete)

		v1.GET("/order", mw, orderController.Index)
		v1.GET("/order/get-by-customer", mw, orderController.GetByCustomer)
		v1.GET("/order/complete", mw, orderController.Complete)
		v1.POST("/order", mw, orderController.Create)
		v1.GET("/order/:id", mw, orderController.GetByID)
		v1.PATCH("/order/:id", mw, orderController.Update)
		v1.DELETE("/order/:id", mw, orderController.Delete)

		v1.GET("/product", mw, productController.Index)
		v1.POST("/product", mw, productController.Create)
		v1.GET("/product/:id", mw, productController.GetByID)
		v1.PATCH("/product/:id", mw, productController.Update)
		v1.DELETE("/product/:id", mw, productController.Delete)
	}
}
