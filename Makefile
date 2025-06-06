# Uber Eats Workflow Assessment Makefile

.PHONY: build clean test run-worker run-trigger run-signal

# Build the Uber Eats workflow
build:
	go build -o bin/uber-eats ./uber-eats

# Clean build artifacts
clean:
	rm -rf bin/

# Test the workflow
test:
	go test ./...

# Run the worker
run-worker:
	./bin/uber-eats -m worker

# Run workflow trigger
run-trigger:
	./bin/uber-eats -m trigger

# Send accept signal (requires WORKFLOW_ID environment variable)
run-signal:
	@if [ -z "$(WORKFLOW_ID)" ]; then \
		echo "Please set WORKFLOW_ID environment variable"; \
		echo "Usage: make run-signal WORKFLOW_ID=your_workflow_id"; \
		exit 1; \
	fi
	./bin/uber-eats -m signal -w $(WORKFLOW_ID)

# Setup local Cadence server
setup-cadence:
	curl -LO https://raw.githubusercontent.com/cadence-workflow/cadence/refs/heads/master/docker/docker-compose.yml
	docker-compose up -d

# Stop Cadence server
stop-cadence:
	docker-compose down

# Complete demo sequence
demo: build
	@echo "🎯 Starting Uber Eats Workflow Demo"
	@echo "1. Building workflow..."
	@echo "2. Ready to run! Execute the following in separate terminals:"
	@echo "   Terminal 1: make run-worker"
	@echo "   Terminal 2: make run-trigger"
	@echo "   Terminal 3: make run-signal WORKFLOW_ID=<id_from_trigger>"
	@echo "   Web UI: http://localhost:8088"

# Help
help:
	@echo "Available commands:"
	@echo "  build         - Build the Uber Eats workflow binary"
	@echo "  clean         - Remove build artifacts"
	@echo "  test          - Run tests"
	@echo "  run-worker    - Start the workflow worker"
	@echo "  run-trigger   - Trigger a new workflow"
	@echo "  run-signal    - Send accept signal (requires WORKFLOW_ID)"
	@echo "  setup-cadence - Start local Cadence server"
	@echo "  stop-cadence  - Stop Cadence server"
	@echo "  demo          - Show demo instructions"
	@echo "  help          - Show this help message"