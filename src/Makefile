# https://www.gnu.org/software/make/manual/html_node/How-Make-Works.html
.DEFAULT_GOAL := run

# Customization Variables
GOAPP_NAME ?= urldl

# Fixed Variables
FOO := BAR

# Goals
build:
	@go build -o ${GOAPP_NAME} .
.PHONY: build

build-container:
	@docker build --no-cache --tag ${GOAPP_NAME}:scratch .
.PHONY: build-container

run:
	@go run main.go --help
.PHONY: run

clean:
	@rm ${GOAPP_NAME}
.PHONY: clean