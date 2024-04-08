# Carpet Bomb
Load test cli app

--url: URL of the service to be tested.

--requests: Total number of requests.

--concurrency: Number of concurrent calls.

--help: help

## Build
```shell
docker build -t carpet-bomb .
```

## Run
```shell
docker run carpet-bomb --url=http://google.com --requests=100 --concurrency=10
```