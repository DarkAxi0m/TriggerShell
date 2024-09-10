# Variables
BINARY_NAME = TriggerShell 
INSTALL_DIR = /usr/local/bin
AUTORUN_FILE = $$HOME/.config/autostart/TriggerShell.desktop

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
	@echo "Checking if $(BINARY_NAME) is running..."
	@if pgrep $(BINARY_NAME) > /dev/null; then \
		echo "$(BINARY_NAME) is running. Stopping it..."; \
		sudo killall $(BINARY_NAME); \
	else \
		echo "$(BINARY_NAME) is not running."; \
	fi
	sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Removing autorun file $(AUTORUN_FILE) as current user..."
	rm -f $(AUTORUN_FILE);

run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

setup:
	@read -p "Enter folder to watch (default: ~/Downloads): " watch_dir; \
	watch_dir=$$(realpath $${watch_dir:-$$HOME/Downloads}); \
	read -p "Enter folder for scripts (default: ./scripts): " scripts_dir; \
	scripts_dir=$$(realpath $${scripts_dir:-./scripts}); \
	mkdir -p $$HOME/.config/autostart; \
	echo "[Desktop Entry]" > $(AUTORUN_FILE); \
	echo "Type=Application" >> $(AUTORUN_FILE); \
	echo "Exec=TriggerShell $$watch_dir $$scripts_dir" >> $(AUTORUN_FILE); \
	echo "Hidden=false" >> $(AUTORUN_FILE); \
	echo "NoDisplay=false" >> $(AUTORUN_FILE); \
	echo "X-GNOME-Autostart-enabled=true" >> $(AUTORUN_FILE); \
	echo "Name[en_AU]=TriggerShell" >>$(AUTORUN_FILE); \
	echo "Name=TriggerShell" >> $(AUTORUN_FILE); \
	echo "Comment[en_AU]=" >> $(AUTORUN_FILE); \
	echo "Comment=" >> $(AUTORUN_FILE); \
	chmod 755 $(AUTORUN_FILE); \
	echo "\nTriggerShell.desktop file created and permissions set in $$HOME/.config/autostart";\
	echo "watch_dir=$$watch_dir"; \
	echo "scripts_dir=$$scripts_dir\n"; \
	read -p "Do you want to start TriggerShell now? (y/n): " start_now; \
	if [ "$$start_now" = "y" ]; then \
		TriggerShell $$watch_dir $$scripts_dir & \
		echo "TriggerShell started with watch_dir=$$watch_dir and scripts_dir=$$scripts_dir"; \
	else \
		echo "TriggerShell setup complete. You can start it manually later."; \
	fi

	
