
GO = go
GOFLAGS = -v


TARGET = app

all: build


build:
	$(GO) build $(GOFLAGS) -o $(TARGET) ./cmd/app

run: 
	$(GO) run $(GOFLAGS) ./cmd/app

rebuild : clear build




clear:
	rm -rf $(TARGET)






format:
	find . -name "*.go" -exec go fmt {} \;

lint:
	golangci-lint run


wrk:
	cd wrk/ 
	wrk -t12 -c1000 -d10s -s wrk/my_script.lua http://localhost:8080/order/b563feb7b2b84b6test


# Фактические цели
.PHONY: all build run clear format rebuild lint wrk
