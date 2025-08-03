# Ping-Pong Service

## Make команды

```sh
make dc-build  # собирает докер-контейнер

make dc-run    # запускает докер-контейнер (не в фоне)

make generate  # генерирует swagger документацию

make run       # запускает бинарник приложения
```

## API

```sh
GET /docs   # swagger документация

GET /about  # выводит информацию о проекте

GET /ping   # выводит "pong"
```