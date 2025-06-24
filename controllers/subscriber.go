package controllers

import (
	"encoding/json"
	"fmt"
	"time"
	
	"github.com/g-s-pai/go-payment-service/initializers"
	"github.com/g-s-pai/go-payment-service/models"
)

func HandleOrderCreated(msgData []byte) error {
	var payment models.Payment
	if err := json.Unmarshal(msgData, &payment); err != nil {
		return err
	}

	fmt.Printf("Processing payment for order: %s\n", payment.OrderID)

	payment.Status = "success"
	payment.CreatedAt = time.Now()

	if err := initializers.DB.Create(&payment).Error; err != nil {
		return err
	}

	// write go code below to update the order status in the orders table
	var order models.Order
	if err := initializers.DB.First(&order, "id = ?", payment.OrderID).Error; err != nil {
		return fmt.Errorf("order not found: %w", err)
	}

	// Update the order status in the database
	if err := initializers.DB.Model(&order).Update("status", "success").Error; err != nil {
		return fmt.Errorf("failed to update order status: %w", err)
	}
	 
	return nil
}
