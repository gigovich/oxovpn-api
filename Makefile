.PHONY: deps
deps:  ## downloads dependencies
	go get -d -v ./cmd/oxovpn-api

.PHONY: build
build: deps ## build oxovpn-api binary
	go build -o ./bin/oxovpn-api ./cmd/oxovpn-api

.PHONY: dev
dev: build ## run development server
	./bin/oxovpn-api serve


# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
