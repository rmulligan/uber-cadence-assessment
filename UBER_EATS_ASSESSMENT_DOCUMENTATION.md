# Uber Developer Advocate Assessment - Uber Eats Workflow Implementation

## Overview
This document details the implementation of a simplified Uber Eats workflow using Cadence, demonstrating distributed workflow orchestration, signal handling, child workflows, and activity execution.

## Architecture Design

### Main Workflow: HandleEatsOrder
The primary workflow manages the complete order lifecycle:

```go
func HandleEatsOrder(ctx workflow.Context, userID string, order Order, restaurantID string) error
```

**Parameters:**
- `userID`: Random UUID identifying the customer
- `order`: Order structure containing ID and content (items array)
- `restaurantID`: Random UUID identifying the restaurant

**Flow Implementation:**
1. **Order Receipt**: Activity prints order received message with full details
2. **Signal Waiting**: Workflow blocks waiting for restaurant acceptance signal
3. **Signal Processing**: Receives "accept" signal to proceed
4. **Preparation Time**: 3-second sleep simulating cooking time
5. **Child Workflow**: Spawns DeliverOrder workflow for delivery
6. **Child Execution**: Waits for delivery completion
7. **Final Message**: Activity prints door delivery confirmation

### Child Workflow: DeliverOrder
Handles the delivery process independently:

```go
func DeliverOrder(ctx workflow.Context, orderID string) error
```

**Implementation:**
- 4-second sleep simulating delivery time
- Activity execution for delivery confirmation
- Independent lifecycle with proper logging

### Activities
Three distinct activities handle message printing:

1. **PrintOrderReceived**: Formats and displays order details
2. **PrintOrderDelivered**: Confirms order delivery in child workflow
3. **PrintOrderAtDoor**: Final confirmation message

## Technical Implementation

### Signal Handling
```go
signalChannel := workflow.GetSignalChannel(ctx, "restaurant-decision")
var decision string
signalChannel.Receive(ctx, &decision)
```

- Uses named signal channel "restaurant-decision"
- Blocks workflow execution until signal received
- Validates signal content before proceeding

### Child Workflow Execution
```go
childWorkflowOptions := workflow.ChildWorkflowOptions{
    WorkflowID: fmt.Sprintf("deliver_order_%s", order.ID),
    ExecutionStartToCloseTimeout: time.Hour,
}
err = workflow.ExecuteChildWorkflow(childCtx, DeliverOrder, order.ID).Get(ctx, nil)
```

- Unique workflow ID per order
- Proper timeout configuration
- Synchronous execution with error handling

### Activity Configuration
```go
activityOptions := workflow.ActivityOptions{
    ScheduleToStartTimeout: time.Minute,
    StartToCloseTimeout:    time.Minute,
}
```

- Consistent timeout policies
- Proper context propagation
- Error handling and logging

## Data Structures

### Order Type
```go
type Order struct {
    ID      string   `json:"id"`
    Content []string `json:"content"`
}
```

Clean data structure supporting:
- Unique order identification
- Variable item quantities
- JSON serialization for persistence

## Deployment & Testing

### Local Environment Setup
1. **Cadence Server**: Docker Compose deployment with Cassandra backend
2. **Port Configuration**: Modified to avoid conflicts (18000, 13000, 8088)
3. **Web UI**: Accessible at `localhost:8088`

### Execution Commands
```bash
# Start worker
./bin/uber-eats -m worker

# Trigger workflow
./bin/uber-eats -m trigger

# Send accept signal
./bin/uber-eats -m signal -w "workflow_id"
```

### Sample Execution
```
Order received: Order ID: abc123, User: user456, Restaurant: rest789, Items: [cheeseburger, diet coke]
→ [Restaurant accepts via signal]
→ [3 second preparation]
→ [Child workflow: 4 second delivery]
Order abc123 delivered!
Your order is in front of your door!
```

## Key Technical Decisions

### 1. Signal Channel Design
- **Choice**: Single channel "restaurant-decision" 
- **Rationale**: Simplifies signal handling while supporting accept/reject logic
- **Alternative**: Separate accept/reject channels would add complexity

### 2. Child Workflow Strategy
- **Choice**: Independent DeliverOrder workflow
- **Rationale**: Separates concerns, enables independent scaling and monitoring
- **Benefit**: Delivery can be tracked separately from order processing

### 3. Activity Granularity  
- **Choice**: Three separate activities for messaging
- **Rationale**: Clear separation of concerns, independent retry policies
- **Benefit**: Granular observability and error handling

### 4. Error Handling
- **Strategy**: Fail-fast with detailed logging
- **Implementation**: Comprehensive error checking at each step
- **Monitoring**: Structured logging for debugging and observability

## Cadence Platform Benefits Demonstrated

### 1. Durability
- Workflow state persisted across failures
- Activities retryable with configurable policies
- Child workflows survive parent failures

### 2. Scalability
- Worker processes can scale independently
- Activities distribute across multiple workers
- Signal handling supports high-throughput scenarios

### 3. Observability
- Complete execution history in Cadence web UI
- Structured logging throughout workflow execution
- Timeline visualization of workflow progress

### 4. Developer Experience
- Type-safe Go APIs
- Clear separation of workflow and activity logic
- Intuitive signal and child workflow patterns

## Learning Process Documentation

### Phase 1: Environment Setup (30 minutes)
- **Challenge**: Docker port conflicts with existing services
- **Solution**: Modified docker-compose.yml to use alternative ports
- **Learning**: Cadence's flexible deployment configuration

### Phase 2: Pattern Study (45 minutes)
- **Approach**: Analyzed existing samples (childworkflow, signalcounter)
- **Key Insights**: Signal channel patterns, child workflow lifecycle
- **Application**: Combined patterns for Uber Eats use case

### Phase 3: Implementation (60 minutes)
- **Structure**: Clean separation into main.go, workflows.go, activities.go
- **Testing**: Iterative build-test cycle with immediate feedback
- **Refinement**: Added comprehensive logging and error handling

### Phase 4: Testing & Validation (30 minutes)
- **Verification**: End-to-end workflow execution
- **UI Validation**: Confirmed workflow visibility in Cadence web interface
- **Signal Testing**: Validated signal handling behavior

## Assessment Process Feedback

### Strengths of the Assessment
1. **Real-world Relevance**: Uber Eats scenario directly relates to Uber's business
2. **Comprehensive Coverage**: Tests multiple Cadence features (signals, child workflows, activities)
3. **Practical Focus**: Emphasizes working implementation over theoretical knowledge
4. **Time Management**: 48-hour window allows for learning and iteration

### Suggested Improvements
1. **Clearer Signal Requirements**: Specify exact signal format and timing expectations
2. **Sample Data**: Provide sample order structures and UUIDs for consistency
3. **Evaluation Criteria**: More detailed rubric for assessment scoring
4. **Advanced Scenarios**: Optional bonus requirements for error handling, retries

### Platform Learning Insights
1. **Cadence Strength**: Excellent developer experience with Go client
2. **Documentation Quality**: Comprehensive samples enable rapid learning
3. **Web UI Value**: Visual workflow tracking essential for debugging
4. **Community Resources**: Active GitHub community with helpful examples

## Conclusion

This implementation successfully demonstrates:
- ✅ Complete Uber Eats workflow with all 7 required steps
- ✅ Proper signal handling for restaurant interaction
- ✅ Child workflow execution for delivery process
- ✅ Activity-based messaging with proper logging
- ✅ Timing controls and error handling
- ✅ Cadence web UI integration for monitoring

The solution showcases Cadence's capabilities for building reliable, observable, and scalable distributed workflows while maintaining clean, maintainable code architecture.