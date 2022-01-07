# Fiber-101

[![Coverage Status](https://coveralls.io/repos/github/panachainy/fiber-101/badge.svg)](https://coveralls.io/github/panachainy/fiber-101)
[![Coverage Status](https://img.shields.io/docker/pulls/panachainy/fiber-101.svg)](https://hub.docker.com/r/panachainy/fiber-101)
[![CodeFactor](https://www.codefactor.io/repository/github/panachainy/fiber-101/badge)](https://www.codefactor.io/repository/github/panachainy/fiber-101)

[GitHub Packages](https://github.com/panachainy/fiber-101/pkgs/container/fiber-101)


## About

Try Fiber framework on 101

## Built With

- Golang
- Fiber v2
- GORM
- PG

## Prerequire

- Go compiler
- GOPATH/bin

    ```sh
    # Golang
    export GOPATH=$HOME/go
    export GOROOT="$(brew --prefix golang)/libexec"
    export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
    ```

- [air](https://github.com/cosmtrek/air)
    * [th](https://www.somkiat.cc/golang-live-reload/)

## Getting started

### Env

```sh
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_USER=postgres
export DATABASE_PASSWORD=1234
export DATABASE_NAME=postgres
export DATABASE_SSLMODE=disable
export DATABASE_TIMEZONE=Asia/Shanghai

# Choose [trace, debug, info, warn, error, fatal, panic] default is debug
export LOG_LEVEL=INFO
```

or

```sh
export DATABASE_DSN="host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
```

### Build

`go install`

`go mod download`

### Clean dependency

`go mod tidy -v` or `make tidy`

### Run

`go run server.go`

### Run + watch

`air`

### Publish

`publish.sh`

## Docker

### Pre-require

#### Database PostgreSQL

Run PostgreSQL

```sh
cd ./mocks/pg/
docker-compose up -d
```

### Docker build

```sh
docker build -t fiber .
```

### Docker try run

```sh
docker run --rm -d \
    --name dev-fiber \
    --network dev-network \
    -p 5050:5050 \
    fiber
```

### Docker run

```sh
docker run \
    --name dev-fiber \
    --network dev-network \
    -p 5050:5050 \
    fiber
```

or

```sh
docker rm fiber
docker run \
    -e DATABASE_HOST=pg_postgres_1 \
    -e DATABASE_PORT=5432 \
    -e DATABASE_USER=postgres \
    -e DATABASE_PASSWORD=1234 \
    -e DATABASE_NAME=postgres \
    -e DATABASE_SSLMODE=disable \
    -e DATABASE_TIMEZONE=Asia/Shanghai \
    --network pg_pg_network \
    -p 5050:5050 \
    fiber
```

## Docker pipeline

TODO: xxx

## Other docs

- [Digital Ocean](docs/digital-ocean.md)
- [Init project](docs/init-project.md)

## Roadmap

See the open [project board](https://github.com/panachainy/fiber-101/projects/1)

## Ref

- [Fiber](https://github.com/gofiber/fiber)
- [Live-reload](https://techinscribed.com/5-ways-to-live-reloading-go-applications/)
- [Intro-golang](https://github.com/AnuchitO/intro-golang)
- [Fiber-Golang-YT](https://www.youtube.com/watch?v=Iq2qT0fRhAA&ab_channel=TutorialEdge)
- [Coveralls](https://github.com/mattn/goveralls)
- [Repository-pattern](https://gramkittisak.medium.com/repository-pattern%E0%B9%83%E0%B8%99-golang-a43411f5479d)
- [Module](https://levelup.gitconnected.com/practical-ddd-in-golang-module-51edf4c319ec)
