.PHONY: build run

DOCKER_IMAGE = go-present
WORK_DIR     = /presentations
CURRENT_PATH = $(shell pwd)

build:
	docker build -t $(DOCKER_IMAGE) .

run: build
	docker run --rm --net="host" --name $(DOCKER_IMAGE) -v $(CURRENT_PATH):$(WORK_DIR) -p 3999:3999 $(DOCKER_IMAGE)
