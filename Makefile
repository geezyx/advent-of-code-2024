# Use PowerShell as the shell on Windows, otherwise use bash
ifeq ($(OS),Windows_NT)
	SHELL := cmd.exe
	.SHELLFLAGS := /C
	RMRF := del /Q /F /S
	MKDIR := mkdir
	NULL := nul
	# Simple zero-padding for Windows using conditional statements
	FMT_DAY = $(if $(filter-out 9,$(N)),0$(N),$(N))
else
	SHELL := /bin/bash
	.SHELLFLAGS := -c
	RMRF := rm -rf
	MKDIR := mkdir -p
	NULL := /dev/null
	FMT_DAY = $(shell printf '%02d' $(N))
endif

.PHONY: help test lint new-day run clean all deps init-day

# Default target when just running 'make'
help:
	@echo "Available targets:"
	@echo "  make test         - Run all tests"
	@echo "  make test N=X     - Run tests for day X"
	@echo "  make lint         - Run linter"
	@echo "  make new-day N=X  - Create new day X (e.g., make new-day N=4)"
	@echo "  make run N=X      - Run solution for day X (e.g., make run N=4)"
	@echo "  make clean        - Remove build artifacts"
	@echo "  make all          - Run tests and linter"
	@echo "  make init-day N=X - Create new day and run initial setup"

# Run all tests
test:
	@echo "Running tests..."
	$(if $(N),\
		go test -v ./day$(FMT_DAY)/...,\
		go test -v ./...)

# Run linter
lint:
	@echo "Running linter..."
	$(if $(shell where golangci-lint 2>$(NULL)),\
		golangci-lint run,\
		(echo "golangci-lint not found. Installing..." && \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
		golangci-lint run))

# Create new day
new-day:
	$(if $(N),,\
		(echo "Please specify day number with N=X (e.g., make new-day N=4)" && \
		exit 1))
	@echo "Creating new day $(N)..."
	go run scripts/newday.go $(N)

# Run specific day
run:
	$(if $(N),,\
		(echo "Please specify day number with N=X (e.g., make run N=4)" && \
		exit 1))
	@echo "Running day $(N)..."
	cd day$(FMT_DAY) && go run .

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	$(if $(OS:Windows_NT=),\
		find . -name "*.test" -type f -delete && \
		find . -name "*.out" -type f -delete,\
		powershell -Command "\
			Get-ChildItem -Recurse -Include *.test,*.out | \
			ForEach-Object { Remove-Item -Force $$_ }")
	go clean

# Run all checks
all: test lint

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	$(if $(shell where golangci-lint 2>$(NULL)),,\
		(echo "Installing golangci-lint..." && \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest))

# Initialize a new day and run tests
init-day: deps new-day test