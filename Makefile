
GO = go
GOFLAGS = -v


TARGET = app

all: build

rebstorage: remegrate migrate

migrate:
	migrate -path migrations -database "postgres://elmir:1902@localhost/wb_L0?sslmode=disable" up

remegrate:
	migrate -path migrations -database "postgres://elmir:1902@localhost/wb_L0?sslmode=disable" down


build:
	$(GO) build $(GOFLAGS) -o $(TARGET) ./cmd/app


run: 
	$(GO) run $(GOFLAGS) ./cmd/app


clear:
	rm -rf $(TARGET)

rebuild : clear build

format:
	find . -name "*.go" -exec go fmt {} \;

# Фактические цели
.PHONY: all build run clear format rebuild migrate remegrate rebstorage
