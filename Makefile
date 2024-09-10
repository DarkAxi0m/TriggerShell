# Variables
BINARY_NAME = TriggerShell 
INSTALL_DIR = /usr/local/bin

# Default target
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

install: 
	@echo "Installing $(BINARY_NAME) to $(INSTALL_DIR)..."
	install -m 0755 $(BINARY_NAME) $(INSTALL_DIR)/$(BINARY_NAME)

uninstall:
	@echo "Uninstalling $(BINARY_NAME) from $(INSTALL_DIR)..."
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)


	
