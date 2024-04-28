#!/bin/sh
ls /migrations
migrate -path=/migrations -database=postgres://wb:wb@postgres:5432/wb?sslmode=disable up

