BINARY_NAME = "go-essentials"
BINARIES = "./bin"
GITHUB = "github.com/max-rodziyevsky/go-essentials"

.PNONY: init
init:
	@echo "::> Creating a module root..."
	@go mod init ${GITHUB}
	@go mod tidy
	@echo "::> Finished!"

.PNONY: build
build:
	@echo "::> Building..."
	@go build -o ${BINARIES}/${BINARY_NAME} ./
	@echo "::> Finished!"

.PNONY: run
run:
	@go build -o ${BINARIES}/${BINARY_NAME} ./
	@${BINARIES}/${BINARY_NAME}

.PNONY: clean
clean:
	@echo "::> Cleaning..."
	@go clean
	@rm -rf ${BINARIES}
	@go mod tidy
	@echo "::> Finished"