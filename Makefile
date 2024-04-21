
GO = go
GOFLAGS = -v


TARGET = app

all: build


build:
	format
	$(GO) build $(GOFLAGS) -o $(TARGET) ./cmd/app


run: 
	format
	$(GO) run $(GOFLAGS) ./cmd/app


clear:
	rm -rf $(TARGET)

rebuild : clear build

format:
	find . -name "*.go" -exec go fmt {} \;

# Фактические цели
.PHONY: all build run clear format rebuild
