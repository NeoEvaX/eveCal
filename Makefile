PACKAGES := $(shell go list ./...)
APP_NAME := 'eveCal'
name := $(shell basename ${PWD})

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## init: initialize project (make init module=github.com/user/project)
.PHONY: init
init:
	go mod init ${module}
	go install github.com/cosmtrek/air@latest
	asdf reshim golang

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## fmt: format code
.PHONY: fmt
fmt:
	gofumpt -l -w .

## test: run unit tests
.PHONY: test
test:
	go test -race -cover $(PACKAGES)

## build: build a binary
.PHONY: build
build: test
	make css && make templ-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

## deploy: build a binary for deployment
.PHONY: deploy
deploy:
	go install github.com/a-h/templ/cmd/templ@latest
	make templ-generate && go build -o ./bin/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

## start: build and run local project
.PHONY: dev
dev:
	go build -o ./tmp/$(APP_NAME) ./cmd/$(APP_NAME)/main.go && air

## css: build tailwindcss
.PHONY: css
css:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch

## templ-generate: generate templates
.PHONY: templ-generate
templ-generate:
	templ generate

## temple-watch: watch generate templates
.PHONY: templ-watch
templ-watch:
	templ generate --watch

## sqlc-generate: generate sqlc files
.PHONY: sqlc-generate
sqlc-generate:
	sqlc generate
