# Тестовый проект с использованием [faktory](https://github.com/contribsys/faktory) для решения фоновых задач
По материалам https://www.youtube.com/watch?v=ZEsLeShY_NY

## Установка
Установка faktory в Docker ([официальная страница установки](https://github.com/contribsys/faktory/wiki/Installation)):
```bash
docker pull contribsys/faktory
```

Запуск контейнера в Docker в РЕЖИМЕ РАЗРАБОТКИ (БД очищается каждый раз):
```bash
docker run --rm -it -p 127.0.0.1:7419:7419 -p 127.0.0.1:7420:7420 contribsys/faktory:latest
```

Запуск контейнера в Docker в ПРОД:
Создать папку Data в желаемом месте, например `data`

```bash
mkdir data
docker run --rm -it \
  -v ./data:/var/lib/faktory/db \
  -e "FAKTORY_PASSWORD=some_password" \
  -p 127.0.0.1:7419:7419 \
  -p 127.0.0.1:7420:7420 \
  contribsys/faktory:latest \
  /faktory -b :7419 -w :7420 -e production
```
В этом случае faktory создаст файл ./data/faktory.rdb для хранения данных фоновых задач. Можно делать бэкапы этого файла.

После запуска на странице http://localhost:7420/ в браузере можно посмотреть дашборд.

## Установка пакета faktory для Go
В проекте выполнить команду:
```bash
go get -u github.com/contribsys/faktory_worker_go
```

## Проверка работоспособности
Запуск воркера:
```bash
go run ./consumer/main.go
```

Запуск продюсера:
```bash
go run ./producer/main.go
```

