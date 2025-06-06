# Uber Developer Advocate Assessment

## Simplified Uber Eats Workflow using Cadence

This repository contains my implementation of the Uber Developer Advocate technical assessment, which involves creating a simplified Uber Eats workflow using the Cadence workflow orchestration platform.

## Assessment Overview

**Objective**: Implement a distributed workflow that simulates the Uber Eats order process, demonstrating proficiency with Cadence workflow patterns including signals, child workflows, activities, and timing controls.

**Technology Stack**: Go, Cadence Workflow Platform, Docker

**Timeline**: 48 hours

## Implementation

### Workflow Architecture

#### Main Workflow: `HandleEatsOrder`
Manages the complete order lifecycle with these steps:
1. **Order Receipt**: Print order received message with details
2. **Signal Waiting**: Wait for restaurant accept/reject signal  
3. **Signal Processing**: Handle restaurant decision
4. **Preparation**: 3-second sleep for cooking simulation
5. **Delivery Initiation**: Spawn child workflow for delivery
6. **Delivery Completion**: Wait for delivery workflow completion
7. **Final Confirmation**: Print door delivery message

#### Child Workflow: `DeliverOrder`
Handles the delivery process independently:
- 4-second delivery simulation
- Order delivery confirmation activity

### Key Features Implemented
- ✅ **Signal Handling**: Restaurant accept/reject workflow signals
- ✅ **Child Workflows**: Independent delivery process management
- ✅ **Activities**: Granular business logic execution with proper timeouts
- ✅ **Timing Controls**: Realistic preparation and delivery simulations
- ✅ **Error Handling**: Comprehensive error checking and logging
- ✅ **Observability**: Structured logging throughout execution

## Project Structure

```
uber-eats/
├── main.go           # Entry point, CLI, and worker setup
├── workflows.go      # HandleEatsOrder and DeliverOrder implementations  
├── activities.go     # Message printing activities
└── go.mod           # Go module dependencies

Documentation/
├── UBER_EATS_ASSESSMENT_DOCUMENTATION.md  # Technical implementation details
├── DEMO_INSTRUCTIONS.md                    # Step-by-step demo guide
└── ASSESSMENT_FEEDBACK.md                  # Process feedback and suggestions
```

## Quick Start

### Prerequisites
- Go 1.22+
- Docker & Docker Compose
- Cadence server running locally

### Setup Instructions

1. **Clone this repository**
   ```bash
   git clone https://github.com/rmulligan/uber-cadence-assessment.git
   cd uber-cadence-assessment
   ```

2. **Start Cadence server**
   ```bash
   # Download docker-compose file
   curl -LO https://raw.githubusercontent.com/cadence-workflow/cadence/refs/heads/master/docker/docker-compose.yml
   docker-compose up -d
   ```

3. **Build the workflow**
   ```bash
   go build -o bin/uber-eats uber-eats/*.go
   ```

4. **Run the demo**
   ```bash
   # Start worker (in one terminal)
   ./bin/uber-eats -m worker
   
   # Trigger workflow (in another terminal)
   ./bin/uber-eats -m trigger
   
   # Send accept signal (replace with actual workflow ID)
   ./bin/uber-eats -m signal -w "workflow_id_from_trigger_output"
   ```

5. **View in Cadence Web UI**
   - Open: http://localhost:8088
   - Navigate to workflows to see execution details

## Demo

For detailed demo instructions, see: [DEMO_INSTRUCTIONS.md](DEMO_INSTRUCTIONS.md)

## Technical Documentation

For comprehensive technical details, see: [UBER_EATS_ASSESSMENT_DOCUMENTATION.md](UBER_EATS_ASSESSMENT_DOCUMENTATION.md)

## Assessment Feedback

For process feedback and improvement suggestions, see: [ASSESSMENT_FEEDBACK.md](ASSESSMENT_FEEDBACK.md)

## Key Learning Outcomes

Through this assessment, I demonstrated:

- **Rapid Technology Adoption**: Learned Cadence from scratch and implemented working solution
- **Distributed Systems Understanding**: Proper workflow orchestration patterns
- **Code Quality**: Clean, maintainable Go code following best practices  
- **Documentation Skills**: Comprehensive technical and process documentation
- **Developer Advocacy Mindset**: Provided constructive feedback for process improvement

## Contact

**Ryan Mulligan**  
Email: ryan@mulligan.dev  
GitHub: [@rmulligan](https://github.com/rmulligan)

---

*This assessment was completed as part of the application process for the Developer Advocate position at Uber, focusing on the Cadence workflow orchestration platform.*