# Uber Eats Workflow Demo Instructions

## Pre-Demo Setup

### 1. Start Cadence Services
```bash
cd /home/ryan/dev/job-search/uber/cadence-samples
docker-compose up -d
```

### 2. Verify Services Running
```bash
docker-compose ps
```
Expected: All services should show "Up" status

### 3. Access Cadence Web UI
Open browser to: `http://localhost:8088`

## Demo Execution

### Step 1: Start the Worker
```bash
./bin/uber-eats -m worker
```

**Demo Points:**
- Shows worker registration with Cadence
- Demonstrates workflow and activity registration
- Logs indicate connection to Cadence server

### Step 2: Trigger Workflow (New Terminal)
```bash
./bin/uber-eats -m trigger
```

**Demo Points:**
- Workflow starts with generated UUIDs
- Note the WorkflowID in output (e.g., `eats_order_4188d031-4561-40ca-adc9-4d465d8c600c`)
- Workflow begins executing and waits for signal

### Step 3: Show Workflow in Web UI
1. Navigate to `http://localhost:8088`
2. Click on "Default" domain
3. Find your workflow (search by WorkflowID if needed)
4. Click to view workflow details

**Demo Points:**
- Show workflow execution timeline
- Point out workflow is "Running" and waiting for signal
- Show input parameters (UserID, Order, RestaurantID)
- Demonstrate workflow history view

### Step 4: Send Restaurant Accept Signal
```bash
./bin/uber-eats -m signal -w "eats_order_4188d031-4561-40ca-adc9-4d465d8c600c"
```
*Replace with actual WorkflowID from Step 2*

**Demo Points:**
- Signal sent to workflow
- Workflow immediately resumes execution
- Watch worker terminal for activity logs

### Step 5: Observe Complete Execution
**In Worker Terminal, you'll see:**
```
Your order received! Order ID: [uuid], User: [uuid], Restaurant: [uuid], Items: [cheeseburger, diet coke]
→ [3 second pause for preparation]
→ [Child workflow starts]
→ [4 second pause for delivery] 
Order [uuid] delivered!
Your order is in front of your door!
```

### Step 6: Review Completed Workflow in UI
Return to web UI and refresh:
- Workflow status now shows "Completed"
- Timeline shows all execution steps
- Child workflow visible as separate execution
- Activity executions visible with timing

## Key Demo Talking Points

### Architecture Highlights
1. **Signal Handling**: Workflow pauses until restaurant accepts
2. **Child Workflows**: Delivery handled as independent process
3. **Activities**: Clear separation of business logic
4. **Timing Controls**: Realistic preparation and delivery times

### Cadence Platform Benefits
1. **Durability**: Workflow state persisted through signals
2. **Observability**: Complete execution history visible
3. **Scalability**: Workers can scale independently
4. **Developer Experience**: Clean Go APIs and patterns

### Implementation Quality
1. **Error Handling**: Comprehensive error checking
2. **Logging**: Structured logging throughout
3. **Configuration**: Proper timeouts and options
4. **Code Structure**: Clean separation of concerns

## Troubleshooting

### If Workflow Doesn't Start
- Check Cadence services: `docker-compose ps`
- Verify worker is running and connected
- Check for port conflicts

### If Signal Doesn't Work
- Ensure correct WorkflowID
- Verify workflow is in "Running" state
- Check worker logs for errors

### If Web UI Not Accessible
- Confirm port 8088 is available
- Check cadence-web service status
- Try refreshing browser

## Demo Cleanup
```bash
# Stop worker (Ctrl+C)
# Stop Cadence services
docker-compose down
```

## Questions to Prepare For

1. **How would you handle order cancellation?**
   - Add cancellation signal channel
   - Implement compensation activities
   - Handle child workflow termination

2. **How would you scale this for high volume?**
   - Multiple worker instances
   - Activity task routing
   - Database partitioning strategies

3. **What about error scenarios?**
   - Retry policies for activities
   - Timeout handling
   - Dead letter queues

4. **How would you monitor this in production?**
   - Metrics collection
   - Alerting on workflow failures
   - Performance monitoring

This demo showcases real-world distributed workflow orchestration and demonstrates readiness for a Developer Advocate role at Uber.