.PHONY: build all

IMAGE=real-life-k8s

all: build
	docker run --rm --net="host" -v `pwd`:/presentations $(IMAGE)

build:
	docker build -t $(IMAGE) .
