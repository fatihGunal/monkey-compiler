BINARY_NAME=monkey
BIN_DIR=bin
CMD_DIR=cmd
MAIN=$(CMD_DIR)/main.go

build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(MAIN)

run: build
	@echo "🚀 Running $(BINARY_NAME)..."
	@./$(BIN_DIR)/$(BINARY_NAME)

test:
	@echo "🧪 Running tests..."
	go test ./...

clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BIN_DIR)

fmt:
	@echo "🎨 Formatting code..."
	go fmt ./...

lint:
	@echo "🔍 Linting code..."
	go vet ./...
