RUN_GO ?= go

test:
	$(RUN_GO) test github.com/...

fmt:
	$(RUN_GO) fmt github.com/...
