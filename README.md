# Golang Fiber MobileApp

# Тестовое задание от компании SMARTIX на языке Java. Я Реализовал весь функционал приложения на языке Golang.
Ссылка на задание: https://docs.google.com/document/d/1PQWv76R9ebtmN5g0CVo1Uvo_DVyCkoZtZMpbfy4QFPA/edit

# Регистрация пользователя /api/v1/register с указанием данных - fname(ФИО), email, phonenumber(логин), password, gender(пол), birthday(дата рождения)
POST запрос на регистрацию
![Alt text](prew/register.png?raw=true "register")

# Отправка платежа на номер с указанием суммы. Если сумма больше баланса или некорректна - ошибка. С баланса списываются денежные средства.
POST запрос на платеж с указанием JWT токена
![Alt text](prew/pay.png?raw=true "pay")
Ошибка при отправке платежа, баланс меньше указанной суммы
![Alt text](prew/payerror.png?raw=true "payerror")
Получение баланса с указанием JWT токена
![Alt text](prew/balanceafterpay.png?raw=true "balance")
Получение истории платежей пользователя с указанием JWT токена
![Alt text](prew/history.png?raw=true "history")
PUT запрос на получение нового токена, если этот просрочился /api/v1/:email/:password/token/new
![Alt text](prew/newtoken.png?raw=true "newtoken")
PUT запрос на обновление данных пользователя с указанием данных fname, email, gender, birthday.

Старт - `docker-compose up` потом - `go run main.go`
Если проблема с PostgreSQL, то нужно переместить create_db.sql на машину командой - `psql -U postgres postgres -h 127.0.0.1 < create_db.sql`
Либо в Docker руками перекинуть в папку и проинициализировать командой - `psql -U postgres postgres < create_db.sql`