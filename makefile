build:
	go build -o bin/pricefetcher

run: build
	./bin/pricefetcher

build_docker:
	docker build --tag fetcher:1 .  
run_docker:
	docker run -p 3000:3000 fetcher:1  