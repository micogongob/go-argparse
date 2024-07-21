RUN_GO ?= docker-compose run --rm go go

test:
	$(RUN_GO) test ./parse

fmt:
	$(RUN_GO) fmt ./...
