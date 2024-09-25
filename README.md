# Golang Fiber MobileApp

# Тестовое задание от компании SMARTIX на языке Java. Я Реализовал весь функционал приложения на языке Golang.
Ссылка на задание: https://docs.google.com/document/d/1PQWv76R9ebtmN5g0CVo1Uvo_DVyCkoZtZMpbfy4QFPA/edit

Старт - `docker-compose up` потом - `go run main.go`

Если проблема с PostgreSQL, то нужно переместить create_db.sql на машину командой - `psql -U postgres postgres -h 127.0.0.1 < create_db.sql`
Либо в Docker руками перекинуть в папку и проинициализировать командой - `psql -U postgres postgres < create_db.sql`