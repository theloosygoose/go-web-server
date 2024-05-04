run:
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@templ generate
	@go run cmd/goserver/main.go

build:
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@npx tailwindcss -o ./dist/tailwind.css --minify
	@templ generate
	@go build -o ./bin/main ./cmd/goserver/main.go

generate:
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@templ generate

update:
	@git pull -f
	@sudo service goweb restart
