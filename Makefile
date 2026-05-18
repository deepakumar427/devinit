# Makefile
.PHONY: build install run clean

# Build binary for current OS/arch
build:
	go build -o bin/devinit ./main.go

# Install to $GOPATH/bin so `devinit` works system-wide
install:
	go install ./...

# Run with example args (useful during development)
run:
	go run ./main.go init my-test-app

# Clean up built binaries
clean:
	rm -rf bin/

# Build for multiple platforms
release:
	GOOS=linux   GOARCH=amd64 go build -o bin/devinit-linux-amd64   ./main.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/devinit-darwin-amd64  ./main.go
	GOOS=windows GOARCH=amd64 go build -o bin/devinit-windows-amd64.exe ./main.go