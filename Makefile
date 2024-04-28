run:
	@templ generate
	@go run cmd/goserver/main.go

build:
	@templ generate
	@go build cmd/goserver/main.go
	@mv ./main ./bin/main

generate:
	@templ generate
