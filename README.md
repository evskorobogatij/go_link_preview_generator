# Service for generate link preview

## Install

For build service you need a **go** installed in your system or run service using [docker](#docker)
```
go mod download
go build -a -installsuffix cgo -ldflags="-w -s" -o build/go_link_preview_generator .
```

## Environment

You can use this environment variables to run service
* `GP_LISTEN` - ip address for run service. Default is `localhost`
* `GP_PORT` - port for run service. Default is `8380`
* `USE_CACHE` - if `1` run service in cache mode
* `REDIS_HOST` - redis host for cacheing
* `REDIS_PORT` - redis post for cacheing

## Usadge

You must send POST request on `/generate_preview` endpoint. Body parameters must contain url - link of site for create preview

## Example 

Just run 
```
# curl --request POST 'http://127.0.0.1:8750/generate_preview' \
> --header 'Content-Type: application/json' \
> --data-raw '{
>     "url": "https://github.com/redis/go-redis"
> }'
```
Anwer is 
```
{
    "title": "GitHub - redis/go-redis: Redis Go client",
    "description": "Redis Go client. Contribute to redis/go-redis development by creating an account on GitHub.",
    "preview_url": "https://opengraph.githubassets.com/3fc8785a3ffcac12c0b4b1d258fb77b24e25bb303852c09444f8e45cadb9014d/redis/go-redis"
}
```

## Docker

You can run this service with Docker. For do this, just type in command prompt:
```
docker build -t go_link_preview_generator -f .\Dockerfile .
# then start container
docker run -it --rm -p 8380:8380 go_link_preview_generator
```

Or yuo can run service this docer composer:
```
docker-compose up
```
With docker-composer servce started with redis and cache renerator result.

In docker composer service started at _8750_ post.