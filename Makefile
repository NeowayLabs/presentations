.PHONY: build present

build:
	docker build -t neoway-presentations .

present: build
	docker run --name neoway-presentations -v `pwd`:/presentations

