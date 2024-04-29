# Используйте официальный образ Go как базовый
FROM golang:1.22-alpine as builder

RUN apk --no-cache add gcc make git bash musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download



COPY cmd/ ./
RUN ls
RUN go build -o ./bin/app ./app/main.go




