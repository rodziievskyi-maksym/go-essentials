BINARY_NAME = "go-essentials"
BINARIES = "./bin"
GITHUB = "github.com/max-rodziyevsky/go-essentials"

init:
	@echo "::> Creating a module root..."
	@go mod init ${GITHUB}
	@go mod tidy
	@echo "::> Finished!"

build:
	@echo "::> Building..."
	@go build -o ${BINARIES}/${BINARY_NAME} ./
	@echo "::> Finished!"

run:
	@go build -o ${BINARIES}/${BINARY_NAME} ./
	@${BINARIES}/${BINARY_NAME}

clean:
	@echo "::> Cleaning..."
	@go clean
	@rm -rf ${BINARIES}
	@go mod tidy
	@echo "::> Finished"

.PNONY: init build run clean