.PHONY: swag run lint 

swag:
	swag init 

lint:
	golangci-lint run -v

run: swag 
	go run main.go
