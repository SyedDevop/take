.PHONY: run

# Define variables
FILE := hello/game/tt.go
RUN_ARGS := go run .

# Default target
run:
	# Print file and path message
	@echo "Getting File and path from: $(FILE)"
	# Display the file and path
	@echo -n "Path: " && $(RUN_ARGS) -d $(FILE) && echo
	@echo -n "File: " && $(RUN_ARGS) -f $(FILE) && echo

install: build
	@mkdir -p ~/app/linux-setup/bin
	@cp -v ./bin/path ~/app/linux-setup/bin/

build:
	@go build -ldflags "-s -w" -o ./bin/path .
