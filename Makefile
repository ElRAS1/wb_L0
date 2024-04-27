
GO = go
GOFLAGS = -v


TARGET = app

all: build


migrate:
	migrate -path migrations -database "postgres://elmir:1902@localhost/wb_L0?sslmode=disable" up

remegrate:
	migrate -path migrations -database "postgres://elmir:1902@localhost/wb_L0?sslmode=disable" down

rebstorage: remegrate migrate





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


# Фактические цели
.PHONY: all build run clear format rebuild migrate remegrate rebstorage lint
