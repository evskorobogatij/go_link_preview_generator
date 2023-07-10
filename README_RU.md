# Сервис для генерации предпросмотра ссылок

## Установка

Для сборки сервиса нуже установленый **go** или можно использовать [docker](#docker)
```
go mod download
go build -a -installsuffix cgo -ldflags="-w -s" -o build/go_link_preview_generator .
```

## Переменные окружения

Сервис использует следующие переменные окружения
* `GP_LISTEN` - ip адрес для запуска сервиса. По-умолчанию `localhost`
* `GP_PORT` - порт на котором будет запущен сервис. Поумолчанию `8380`
* `USE_CACHE` - признак кеширования. Для того чтобы кеширование было доступно нужно установить в `1`
* `REDIS_HOST` - адрес redis-сервиса для кеширования
* `REDIS_PORT` - порт redis-сервиса для кеширования

## Использование

Необходимо отправить POST запрос на `/generate_preview`. В теле запроса url - ссылка на сайт для создания предпросмотра

## Пример 

Запустите 
```
# curl --request POST 'http://127.0.0.1:8750/generate_preview' \
> --header 'Content-Type: application/json' \
> --data-raw '{
>     "url": "https://github.com/redis/go-redis"
> }'
```
Ответ 
```
{
    "title": "GitHub - redis/go-redis: Redis Go client",
    "description": "Redis Go client. Contribute to redis/go-redis development by creating an account on GitHub.",
    "preview_url": "https://opengraph.githubassets.com/3fc8785a3ffcac12c0b4b1d258fb77b24e25bb303852c09444f8e45cadb9014d/redis/go-redis"
}
```

## Docker

Для засуска сервиса можно использовать Docker. Для этого введите в консоли:
```
docker build -t go_link_preview_generator -f .\Dockerfile .
# then start container
docker run -it --rm -p 8380:8380 go_link_preview_generator
```

Также можно воспользоваться docker-composer:
```
docker-compose up
```
При использовании docker-composer сервис запускается вместе с redis и включенным режимом кеширования.

В docker-composer сервис запускается на порту _8750_.