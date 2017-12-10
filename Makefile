.PHONY: install
install: up
	npm --prefix ./www i

.PHONY: up
up: build
	docker-compose up -d

.PHONY: build
build:
	docker-compose build

.PHONY: stop
stop:
	docker-compose stop

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: deps
deps:
	docker-compose exec app go get ./...

.PHONY: test
test:
	go test -v ./src

.PHONY: docker-test-acceptance
docker-test-acceptance:
	docker-compose exec app go test -v -tags=acceptance

.PHONY: client-watch
client-watch:
	npm --prefix ./www run test

.PHONY: client-test
client-test:
	npm --prefix ./www run test-single-run

.PHONY: full-test
full-test: docker-test-acceptance
	npm --prefix ./www run test-single-run
