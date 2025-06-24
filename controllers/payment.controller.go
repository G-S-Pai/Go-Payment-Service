package controllers

import (
	"github.com/g-s-pai/go-payment-service/initializers"
	"github.com/g-s-pai/go-payment-service/models"

	"github.com/kataras/iris/v12"
)

func GetPayments(ctx iris.Context) {
	var payments []models.Payment
	if err := initializers.DB.Find(&payments).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.JSON(payments)
}

func GetPaymentByOrderID(ctx iris.Context) {
	orderId := ctx.Params().Get("id")
	var payment models.Payment
	if err := initializers.DB.First(&payment, "order_id = ?", orderId).Error; err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Payment not found"})
		return
	}
	ctx.JSON(payment)
}