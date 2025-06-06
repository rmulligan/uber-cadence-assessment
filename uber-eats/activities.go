package main

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/cadence/activity"
	"go.uber.org/zap"
)

// PrintOrderReceived prints the order received message with details
func PrintOrderReceived(ctx context.Context, userID string, order Order, restaurantID string) error {
	logger := activity.GetLogger(ctx)
	
	orderDetails := strings.Join(order.Content, ", ")
	message := fmt.Sprintf("Your order received! Order ID: %s, User: %s, Restaurant: %s, Items: [%s]",
		order.ID, userID, restaurantID, orderDetails)
	
	fmt.Println(message)
	logger.Info("Order received message printed", 
		zap.String("Message", message),
		zap.String("OrderID", order.ID))
	
	return nil
}

// PrintOrderDelivered prints the order delivered message
func PrintOrderDelivered(ctx context.Context, orderID string) error {
	logger := activity.GetLogger(ctx)
	
	message := fmt.Sprintf("Order %s delivered!", orderID)
	
	fmt.Println(message)
	logger.Info("Order delivered message printed", 
		zap.String("Message", message),
		zap.String("OrderID", orderID))
	
	return nil
}

// PrintOrderAtDoor prints the final delivery confirmation message
func PrintOrderAtDoor(ctx context.Context, orderID string) error {
	logger := activity.GetLogger(ctx)
	
	message := "Your order is in front of your door!"
	
	fmt.Println(message)
	logger.Info("Final delivery message printed", 
		zap.String("Message", message),
		zap.String("OrderID", orderID))
	
	return nil
}