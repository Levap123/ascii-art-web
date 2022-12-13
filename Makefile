run:
	docker run -p 8080:8080 ascii

build:
	docker build -t ascii .	
	
go:
	go run cmd/main.go