.PHONY: build

build :
	go build -v ./cmd/app

.PHONY: run
run:
	go run -v ./cmd/app 
# -config-path configs/app.toml

.PHONY: clear
clear :
	rm -rf app


