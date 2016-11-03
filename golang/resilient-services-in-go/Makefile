.PHONY: build all

IMAGE=resilient-services

all: build
	docker run --rm --net="host" -v `pwd`:/presentations $(IMAGE)

build:
	docker build -t $(IMAGE) .
