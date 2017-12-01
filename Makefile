.PHONY: up
up:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop
