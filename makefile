MAIN = "cmd/main.go"

run:
	go run $(MAIN)

generate:
	swag init -g $(MAIN) --output docs