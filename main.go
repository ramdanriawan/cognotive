package main

import (
	"crypto/tls"
	"encoding/csv"
	"fmt"
	"os"

	config "skyshi.com/src/config"
	route "skyshi.com/src/routes"

	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/go-gomail/gomail"

	// customer "skyshi.com/src/entities/customer"
	"skyshi.com/src/entities/customer"
	order "skyshi.com/src/entities/order"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests. Try again in "+time.Until(info.ResetTime).String())
}

func main() {
	r := gin.Default()

	db := config.DB()

	// This makes it so each ip can only make 5 requests per second
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Minute,
		Limit: 100,
	})
	mw := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	route.Api(r, db, mw)

	s := gocron.NewScheduler(time.UTC)

	s.Cron("0 0 * * *").Do(func() {
		d := gomail.NewDialer("smtp.gmail.com", 587, "ramdanriawan3@gmail.com", "pevkybszideqpout")
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		var orderModel []order.OrderModel

		db.Find(&orderModel, order.OrderModel{Status: "Pending"})

		rows := [][]string{
			{"Order Id", "Customer Name", "Order Date", "Total", "Status"},
		}

		for i := 0; i < len(orderModel); i++ {
			var customerModel customer.CustomerModel
			db.Preload("Orders.OrderDetails.Product").Find(&customerModel, customer.CustomerModel{ID: orderModel[i].CustomerId})

			m := gomail.NewMessage()
			m.SetHeader("From", "ramdanriawan3@gmail.com")
			m.SetHeader("To", customerModel.Email)
			m.SetHeader("Subject", "You have "+fmt.Sprint(len(customerModel.Orders))+" Pending Orders")

			var body = ""
			var total = 0
			for j := 0; j < len(customerModel.Orders); j++ {

				for k := 0; k < len(customerModel.Orders[j].OrderDetails); k++ {
					body += "OrderId: " + fmt.Sprint(customerModel.Orders[j].ID) + "<br>"
					body += "Product Name: " + fmt.Sprint(customerModel.Orders[j].OrderDetails[k].Product.Name) + "<br>"
					body += "Price: " + fmt.Sprint(customerModel.Orders[j].OrderDetails[k].Price) + "<br>"
					body += "Qty: " + fmt.Sprint(customerModel.Orders[j].OrderDetails[k].Qty) + "<br>"
					body += "Total: " + fmt.Sprint(customerModel.Orders[j].OrderDetails[k].Total) + "<br>"
					body += "----------<br>"

					total += customerModel.Orders[j].OrderDetails[k].Total
				}

				rows = append(rows, []string{fmt.Sprint(customerModel.Orders[j].ID), customerModel.Name, fmt.Sprint(customerModel.Orders[j].Date), fmt.Sprint(total), fmt.Sprint(customerModel.Orders[j].Status)})
				total = 0
			}

			body += "<br> <h3>Please click this link to complete the orders: http://localhost:3030/order/complete?user_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MX0.fhc3wykrAnRpcKApKhXiahxaOe8PSHatad31NuIZ0Zg</h3>"

			m.SetBody("text/html", body)

			// Send the email to Bob, Cora and Dan.
			if err := d.DialAndSend(m); err != nil {
				panic(err)
			}

			// untuk export ke csv
			csvfile, err := os.Create("data.csv")

			if err != nil {
				panic(err)
			}

			cswriter := csv.NewWriter(csvfile)

			for _, row := range rows {
				_ = cswriter.Write(row)
			}

			cswriter.Flush()
			csvfile.Close()
		}
	})

	s.StartAsync()

	r.Run(":3030")
}
