build:
	go build  main.go

run:
	go run main.go

pre:
	docker build  -t gin_demo:pre -f deploy/pre/Dockerfile .
test:
	docker build  -t gin_demo:test -f deploy/test/Dockerfile .

