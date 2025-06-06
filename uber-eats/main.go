package main

import (
	"flag"
	"time"

	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"

	"github.com/uber-common/cadence-samples/cmd/samples/common"
)

// ApplicationName is the task list for this sample
const ApplicationName = "uber-eats-workflow"

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *common.SampleHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.WorkerMetricScope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func startWorkflow(h *common.SampleHelper) {
	// Sample order data
	order := Order{
		ID:      uuid.New(),
		Content: []string{"cheeseburger", "diet coke"},
	}
	
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "eats_order_" + uuid.New(),
		TaskList:                        ApplicationName,
		ExecutionStartToCloseTimeout:    time.Hour,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	
	h.StartWorkflow(workflowOptions, HandleEatsOrder, uuid.New(), order, uuid.New())
}

func sendAcceptSignal(h *common.SampleHelper, workflowID string) {
	h.SignalWorkflow(workflowID, "restaurant-decision", "accept")
}

func main() {
	var mode string
	var workflowID string
	flag.StringVar(&mode, "m", "trigger", "Mode is worker, trigger, or signal.")
	flag.StringVar(&workflowID, "w", "", "the workflowID to send accept signal")
	flag.Parse()

	var h common.SampleHelper
	h.SetupServiceConfig()

	switch mode {
	case "worker":
		h.RegisterWorkflow(HandleEatsOrder)
		h.RegisterWorkflow(DeliverOrder)
		h.RegisterActivity(PrintOrderReceived)
		h.RegisterActivity(PrintOrderDelivered)
		h.RegisterActivity(PrintOrderAtDoor)
		startWorkers(&h)

		// The workers are supposed to be long running process that should not exit.
		// Use select{} to block indefinitely for samples, you can quit by CMD+C.
		select {}
	case "trigger":
		startWorkflow(&h)
	case "signal":
		if workflowID == "" {
			panic("workflowID is required for signal mode")
		}
		sendAcceptSignal(&h, workflowID)
	}
}