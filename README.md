# Carpet Bomb
Load test cli app

--url: URL do serviço a ser testado.

--requests: Número total de requests.

--concurrency: Número de chamadas simultâneas.

## Build
```shell
docker build -t carpet-bomb .
```

## Run
```shell
docker run carpet-bomb --url=http://google.com --requests=100 --concurrency=10
```