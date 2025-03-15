### CONFIG_PATH
```bash
export CONFIG_PATH=/home/sj_shoff/url-shortener/config/local.yaml
```

### Очистка docker-compose (вместе с кэшем)
```bash
docker-compose down -v --rmi all
docker system prune -a --volumes

docker-compose down -v && docker-compose up --build
```

### Работа с докером:
```bash
docker run --name=todo-db -e POSTGRES_PASSWORD='03032006' -p 5432:5432 -d --rm postgres
docker ps
```

### Проверка БД в контейнере(1я строка - берем айдишник postgres)
```bash
docker ps
docker exec -it 6188ec4b2fdd /bin/bash
psql -U postgres
\d
```

### Скрипт wait-for-postgres.sh гарантирует, что приложение запустится только после готовности базы данных.

### Проверка занятости порта / освобождение порта
```bash
sudo lsof -i :5432
sudo kill 474
```

### Права для работы с папкой
```bash
sudo chmod -R 777 .database
```

### Очистка docker-compose (вместе с кэшем)
```bash
docker-compose down -v --rmi all
docker system prune -a --volumes
```

### Отключение Firewall
```bash
sudo ufw disable
```

### Подключение к базе данных внутри контейнера Docker и изменение пароля
```bash
docker-compose exec db psql -U postgres -d postgres
ALTER USER postgres WITH PASSWORD '03032006';
\q
```