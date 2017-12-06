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

.PHONY: full-test
full-test:
	go test -v ./src
	npm --prefix ./www run test-single-run
