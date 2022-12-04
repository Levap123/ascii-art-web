run:
	docker run -p 0.0.0.0:80:8080 ascii

build:
	docker build -t ascii .	
	
go:
	go run cmd/main.go