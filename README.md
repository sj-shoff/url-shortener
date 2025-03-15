# URL Shortener - сервис сокращения URL на Go

## В проекте реализовано:
- Разработка Веб-Приложения на Go, следуя дизайну REST API.
- Работа с фреймворком <a href="https://github.com/go-chi/chi">go-chi/chi</a>.
- Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
- Работа с БД Postgres. Запуск из Docker. Docker-compose. Makefile.
- Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Сокращение длинных URL в короткие алиасы
- Редирект по коротким ссылкам
- Базовое управление URL (создание/удаление)
- JWT-like аутентификация через Basic Auth (в будущем переход на gRPC-Authorization-Service)
- Подробное логирование с разными уровнями
- Конфигурация через environment variables
- Graceful shutdown сервера
- Middleware для обработки запросов

### Для запуска приложения:

```
make build && make run
```
