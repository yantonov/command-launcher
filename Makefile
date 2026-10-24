.PHONY: help help-all build install test integration doc-lint check

.DEFAULT_GOAL := help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "  check    - Full verification (build + test + integration)"
	@echo "  build    - Build the binary"
	@echo "  install  - Build and copy binary to \$$CDT_HOME"
	@echo "  test     - Run unit tests"
	@echo ""
	@echo "Run 'make help-all' for the full list, including internal targets."

help-all:
	@echo "Public targets:"
	@echo "  check        - Full verification (build + test + integration)"
	@echo "  build        - Build the binary"
	@echo "  install      - Build and copy binary to \$$CDT_HOME"
	@echo "  test         - Run unit tests"
	@echo ""
	@echo "Internal targets (typically called via check):"
	@echo "  integration  - Run integration tests"
	@echo "  doc-lint     - Lint documentation"

CDT_HOME ?=

install: build
ifeq ($(CDT_HOME),)
	$(error CDT_HOME is not set)
endif
	cp cdt "$(CDT_HOME)/cdt"

build:
	go build -o cdt -ldflags="-X main.version=dev -X main.buildNum=$$(date +'%Y%m%d-%H%M%S') -X main.appName=cdt -X 'main.appLongName=Criteo Dev Toolkit'"

test:
	go test -v ./...

integration:
	./test/integration.sh

doc-lint:
	cd gh-pages && npm install && npm run test

check: build test
	./test/integration.sh
	@echo "ALL CHECKS PASSED"
