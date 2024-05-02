run:
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@templ generate
	@go run cmd/goserver/main.go

build:
	@git pull -f
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@npx tailwindcss -o ./dist/tailwind.css --minify
	@templ generate
	@go build cmd/goserver/main.go
	@mv ./main ./bin/main
	@sudo service goweb restart

generate:
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@templ generate

