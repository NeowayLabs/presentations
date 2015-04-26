.PHONY: build present

build:
	docker build -t neoway-presentations .

present: build
	docker run --rm --net="host" --name neoway-presentations -v `pwd`:/presentations -p 3999:3999 neoway-presentations
