Сервер для динамически управлять конфигурацией приложений

<!-- TOC -->
    * [Пример использования сервиса](#--)
      * [Получение конфига](#-)
      * [Создание конфига](#-)
    * [Запуск сервиса:](http://localhost:2000)
<!-- TOC -->

Сервер реализован на языке GoLang и использует в качестве хранилища базу данных PostgreSQL. 
Конфигурации хранятся в формате JSON.
Поддерживается версионирование для каждого сервиса.
Доступ к серверу осуществляется через REST API

Примеры запросов находятся в query.http

### Пример использования сервиса
#### Получение конфига

`curl http://localhost:8080/config?service=managed-k8s`

```json
{"key1": "value1", "key2": "value2"}
```

#### Создание конфига

`curl -d "@data.json" -H "Content-Type: application/json" -X POST http://localhost:8080/config`

```json
{
    "service": "managed-k8s",
    "data": [
        {"key1": "value1"},
        {"key2": "value2"}
    ]
}
```

PATCH http://localhost:8080/config

Content-Type: application/json

```json

{
  "service": "managed-k8s",
  "data": [
    {"key1": "value101"},
    {"key2": "value202"}
  ]
}
```


```json
curl -X DELETE --location "http://localhost:8080/config?service=managed-k8s" \
    -H "Content-Type: application/json"
```
### Запуск сервиса:
```git clone git@github.com:chelnik/sberCloudServer.git```

```docker-compose up -d```