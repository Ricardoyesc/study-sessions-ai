.PHONY: build run test dev

build:
	$(MAKE) -C src build

run:
	$(MAKE) -C src run

test:
	$(MAKE) -C src test

dev:
	$(MAKE) -C src dev

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker-compose build
