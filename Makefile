run:
	@npx tailwindcss -i ./internal/view/input.css -o ./internal/view/output.css
	@templ generate
	@go run cmd/goserver/main.go

build:
	@npx tailwindcss -i ./internal/view/input.css -o ./internal/view/output.css
	@npx tailwindcss -o ./internal/view/output.css --minify
	@templ generate
	@go build cmd/goserver/main.go
	@mv ./main ./bin/main

generate:
	@npx tailwindcss -i ./internal/view/input.css -o ./internal/view/output.css
	@templ generate
