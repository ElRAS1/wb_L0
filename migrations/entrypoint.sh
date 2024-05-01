#!/bin/sh
migrate -path=/migrations -database=postgresql://wb:wb@postgres:5432/wb?sslmode=disable up

