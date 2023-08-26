# Set the compiler to use for building the program
CC=go

GOFLAGS=-ldflags="-s -w"
OUTPUT=samb-be

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
# Define the build target
build:
	$(CC) build  -o $(OUTPUT) main.go

# Define the clean target
clean:
	rm -f $(OUTPUT)

swag:
	swag init -g main.go

GO := @go
PROTOC := @protoc

protoc:
	${PROTOC} --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/tasklist/*.proto
