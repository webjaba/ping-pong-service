MAIN = "cmd/main.go"

run:
	go run $(MAIN)

generate:
	swag init -g $(MAIN) --output docs

dc-build:
	docker build -t ping-pong-service:v1 .

make dc-run:
	docker run -p 9000:9000 ping-pong-service:v1