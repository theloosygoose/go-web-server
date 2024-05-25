run:
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@templ generate
	@go run cmd/goserver/main.go

build:
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@templ generate
	@go build -o ./bin/main ./cmd/goserver/main.go

generate:
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css
	@templ generate

prod:
	@git pull -f
	@npx tailwindcss -i ./dist/input.css -o ./dist/tailwind.css -m
	@templ generate
	@go build -o ./bin/main ./cmd/goserver/main.go
	@sudo service goweb restart

migrate_up:
	@migrate -path database/migrations/ -database "sqlite3:///mnt/usb/sqlite/donphoto.db" up

migrate_down:
	@migrate -path database/migrations/ -database "sqlite3:///mnt/usb/sqlite/donphoto.db" down 

sqlc:
	@sqlc generate
