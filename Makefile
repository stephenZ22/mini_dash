# 输出目录
BIN_DIR := bin

# 可执行文件名
SERVER_BIN := $(BIN_DIR)/minidash-server
CLIENT_BIN := $(BIN_DIR)/minidash-client

# 默认构建所有
all: build-server build-client

# 构建 server
build-server:
	@echo "Building server..."
	@mkdir -p $(BIN_DIR)
	go build -o $(SERVER_BIN) ./cmd/server

# 构建 client
build-client:
	@echo "Building client..."
	@mkdir -p $(BIN_DIR)
	go build -o $(CLIENT_BIN) ./cmd/client

# 清理编译产物
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

# 运行 server
run-server: build-server
	@$(SERVER_BIN)

# 运行 client
run-client: build-client
	@$(CLIENT_BIN)
