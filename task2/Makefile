#!/usr/bin/make
# Makefile readme (ru): <http://linux.yaroslavl.ru/docs/prog/gnu_make_3-79_russian_manual.html>
# Makefile readme (en): <https://www.gnu.org/software/make/manual/html_node/index.html#SEC_Contents>

.PHONY: run build lint
.DEFAULT_GOAL: help

run: ## Run app
	go run main.go

build: ## Build app
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./build/main main.go

lint: ## Linter run
	 golangci-lint run