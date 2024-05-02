FROM golang:1.22-alpine as builder

RUN apk --no-cache add gcc make git bash musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download



COPY  . /app
RUN make build


FROM alpine


COPY --from=builder /app  /

CMD ["./app"]












