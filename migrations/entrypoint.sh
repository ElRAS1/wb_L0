#!/bin/sh
ls /migrations
# migrate -path=/migrations -database=postgresql://wb:wb@postgres:5432/wb?sslmode=disable up
# migrate -path=/migrations -database=postgresql://wb:wb@postgres:5432/wb?sslmode=disable down --all
migrate -path=/migrations -database=postgresql://wb:wb@postgres:5432/wb?sslmode=disable up

