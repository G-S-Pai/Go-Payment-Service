package main

import (
	"log"
	"github.com/g-s-pai/go-payment-service/initializers"
	"github.com/g-s-pai/go-payment-service/pubsub"
	"github.com/g-s-pai/go-payment-service/routes"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

var (
	app *iris.Application
)

func init() {
	// Load .env variables
	_ = godotenv.Load()
	
	// Connect to DB
	initializers.ConnectDB()
	
	
	// Start listening to Pub/Sub in a goroutine (non-blocking)
	go func() {
		if err := pubsub.ListenForOrders(); err != nil {
			log.Fatalf("Pub/Sub listen error: %v", err)
		}
	}()
		
	// Start HTTP server for health check or future use
	app = iris.New()
}
func main() {
	
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Payment Service is running")
	})

	router := app.Party("/api/v1")

	routes.PaymentRoutes(router)

	if err := app.Listen(":3001"); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
