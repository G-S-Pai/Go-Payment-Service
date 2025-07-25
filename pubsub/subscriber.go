package pubsub

import (
	"context"
	"fmt"
	"os"
	
	"cloud.google.com/go/pubsub"
	"github.com/g-s-pai/go-payment-service/controllers"
)

func ListenForOrders() error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, os.Getenv("PUBSUB_PROJECT_ID"))
	if err != nil {
		return err
	}

	sub := client.Subscription(os.Getenv("ORDERS_SUBSCRIPTION"))

	fmt.Println("Listening for order-created events...")

	return sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))

		if err := controllers.HandleOrderCreated(msg.Data); err != nil {
			fmt.Printf("Failed to process payment: %v\n", err)
			msg.Nack()
		} else {
			msg.Ack()
		}
	})
}
