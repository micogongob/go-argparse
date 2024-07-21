RUN_GO ?= docker-compose run --rm go go

test:
	$(RUN_GO) test ./lib/...

fmt:
	$(RUN_GO) fmt ./lib/...
