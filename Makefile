.PHONY: up
up:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop

.PHONY: build
build:
	docker-compose build

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: deps
deps:
	docker-compose exec app go get ./...
