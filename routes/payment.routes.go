package routes

import (
	"github.com/g-s-pai/go-payment-service/controllers"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
)

func PaymentRoutes(rg router.Party) {
	router := rg.Party("/payments")

	router.Use(iris.Compression)
	router.Get("/", controllers.GetPayments)
	router.Get("/order/{id}", controllers.GetPaymentByOrderID)
}
