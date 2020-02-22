#!/bin/bash
build:
	go get github.com/mattn/go-sqlite3
	go get ./delivery/...
	GOARCH=amd64 GOOS=linux go build -o main ./delivery/main.go
run:
	./main 