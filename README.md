# Template API
### _Шаблон для создания RESTful API_

## Проект создан с использованием языка [Go](https://go.dev/) версии 1.22
## Основные используемые библиотеки:

- [Echo](https://echo.labstack.com) High performance, extensible, minimalist Go web framework
- [GORM](https://gorm.io) The fantastic ORM library for Golang
- [swag](https://github.com/swaggo/swag) Automatically generate RESTful API documentation with Swagger 2.0 for Go

Проект также использует [PostgreSQL](https://www.postgresql.org/) в качестве СУБД

## Запуск

Конфигурация проекта должна находиться в файле `.env`. Для примера в проекте присутствует файл конфигурации `.env.dist`, можно скопировать его содержание и настроить под свое окружение

_пример файла `.env`:_

```
ECHO_HOST=:8020
SWAGGER_HOST=0.0.0.0:8020
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=template
DB_PORT=5432
SSL_MODE=disable
SSL_ROOT_CERTIFICATE=root.crt
```

Для работы проекта потребуется связь с сервером PostgreSQL. Можно используя [Docker](https://www.docker.com/) создать соответствующий контейнер:
```sh
docker run --name postgres_db  -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=template -d postgres
```

Для компиляции и запуска проекта выполните:

```sh
go run cmd/main.go
```

В случае использования переменных среды из файла `.env.dist` доступ к swag можно получить в браузере по адресу
```
0.0.0.0:8020/swagger/
```

## Swagger
Для генерации файлов swagger из корня проекта необходимо запустить:
```sh
swag init -g ./cmd/main.go -o ./docs
```

## Тестирование
Примеры тестов доступны в директории `tests`. Запустить из корня проекта можно командой.

```sh
go test ./tests -v
```
