package main

import (
	"fmt"
	"time"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// Order represents the cart structure
type Order struct {
	ID      string   `json:"id"`
	Content []string `json:"content"`
}

// HandleEatsOrder implements the main Uber Eats workflow
func HandleEatsOrder(ctx workflow.Context, userID string, order Order, restaurantID string) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting HandleEatsOrder workflow",
		zap.String("UserID", userID),
		zap.String("OrderID", order.ID),
		zap.String("RestaurantID", restaurantID),
		zap.Strings("OrderContent", order.Content))

	// Activity options for all activities
	activityOptions := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// Step 1: Print order received message
	err := workflow.ExecuteActivity(ctx, PrintOrderReceived, userID, order, restaurantID).Get(ctx, nil)
	if err != nil {
		logger.Error("Failed to print order received message", zap.Error(err))
		return err
	}

	// Step 2: Wait for restaurant signal to accept or reject
	logger.Info("Waiting for restaurant to accept or reject the order...")
	signalChannel := workflow.GetSignalChannel(ctx, "restaurant-decision")
	var decision string
	signalChannel.Receive(ctx, &decision)
	
	logger.Info("Received restaurant decision", zap.String("Decision", decision))
	
	if decision != "accept" {
		logger.Info("Order was rejected by restaurant")
		return fmt.Errorf("order rejected by restaurant")
	}

	// Step 3: Send accept signal (already received above)
	logger.Info("Order accepted by restaurant!")

	// Step 4: Sleep for 3 seconds (preparation time)
	logger.Info("Order is being prepared...")
	err = workflow.Sleep(ctx, 3*time.Second)
	if err != nil {
		logger.Error("Failed to sleep for preparation", zap.Error(err))
		return err
	}

	// Step 5 & 6: Create and call child workflow DeliverOrder
	logger.Info("Starting delivery process...")
	childWorkflowOptions := workflow.ChildWorkflowOptions{
		WorkflowID:                   fmt.Sprintf("deliver_order_%s", order.ID),
		ExecutionStartToCloseTimeout: time.Hour,
	}
	childCtx := workflow.WithChildOptions(ctx, childWorkflowOptions)
	
	err = workflow.ExecuteChildWorkflow(childCtx, DeliverOrder, order.ID).Get(ctx, nil)
	if err != nil {
		logger.Error("Child delivery workflow failed", zap.Error(err))
		return err
	}

	// Step 7: Print final delivery message
	err = workflow.ExecuteActivity(ctx, PrintOrderAtDoor, order.ID).Get(ctx, nil)
	if err != nil {
		logger.Error("Failed to print final delivery message", zap.Error(err))
		return err
	}

	logger.Info("HandleEatsOrder workflow completed successfully")
	return nil
}

// DeliverOrder implements the child delivery workflow
func DeliverOrder(ctx workflow.Context, orderID string) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting DeliverOrder child workflow", zap.String("OrderID", orderID))

	// Activity options
	activityOptions := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	// Sleep for 4 seconds (delivery time)
	logger.Info("Driver is on the way...")
	err := workflow.Sleep(ctx, 4*time.Second)
	if err != nil {
		logger.Error("Failed to sleep for delivery", zap.Error(err))
		return err
	}

	// Print order delivered message
	err = workflow.ExecuteActivity(ctx, PrintOrderDelivered, orderID).Get(ctx, nil)
	if err != nil {
		logger.Error("Failed to print order delivered message", zap.Error(err))
		return err
	}

	logger.Info("DeliverOrder child workflow completed successfully")
	return nil
}