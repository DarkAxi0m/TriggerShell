# Makefile

build:
	go build -o TriggerShell

clean:
	@echo "Cleaning up..."
	rm -f TriggerShell

run: build
	@echo "Running the program..."
	./TriggerShell
	
