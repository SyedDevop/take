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
	@echo "Install...."
	@mkdir -p ~/app/linux-setup/bin
	@echo "Copying executable to local"
	@cp -v ./bin/path ~/.local/bin/g_path
	@echo "Done .."

build:
	@echo "Building the Executable...."
	@go build -ldflags "-s -w" -o ./bin/path .
