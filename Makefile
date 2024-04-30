run:
	@npx tailwindcss -i ./internal/view/dist/input.css -o ./internal/view/dist/tailwind.css
	@templ generate
	@go run cmd/goserver/main.go

build:
	@npx tailwindcss -i ./internal/view/dist/input.css -o ./internal/view/dist/tailwind.css
	@npx tailwindcss -o ./internal/view/dist/output.css --minify
	@templ generate
	@go build cmd/goserver/main.go
	@mv ./main ./bin/main

generate:
	@npx tailwindcss -i ./internal/view/dist/input.css -o ./internal/view/dist/tailwind.css
	@templ generate
