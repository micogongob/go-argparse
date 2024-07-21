RUN_GO ?= docker-compose run --rm go go

test:
	$(RUN_GO) test ./...

fmt:
	$(RUN_GO) fmt ./...

ci:
	docker login -u $$DOCKER_HUB_USER -p $$DOCKER_HUB_PASS
	docker-compose pull
