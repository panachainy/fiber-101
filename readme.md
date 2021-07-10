# Fiber-101

[![Coverage Status](https://coveralls.io/repos/github/panachainy/fiber-101/badge.svg)](https://coveralls.io/github/panachainy/fiber-101)

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

## Initial project

### Create mod

```sh
go mod init demo
```

## Getting started

### Env

```sh
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=1234
export DB_NAME=postgres
export DB_SSLMODE=disable
export DB_TIMEZONE=Asia/Shanghai
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
    -p 5000:5000 \
    fiber
```

### Docker run

```sh
docker run \
    --name dev-fiber \
    --network dev-network \
    -p 5000:5000 \
    fiber
```

or

> Note: After -e is key of env from your machine.

```sh
docker rm fiber
docker run \
    -e DB_HOST=pg_postgres_1 \
    -e DB_PORT=5432 \
    -e DB_USER=postgres \
    -e DB_PASSWORD=1234 \
    -e DB_NAME=postgres \
    -e DB_SSLMODE=disable \
    -e DB_TIMEZONE=Asia/Shanghai \
    --network pg_pg_network \
    -p 5000:5000 \
    fiber
```

## Other docs

- [Digital Ocean](docs/digital-ocean.md)

## Roadmap

See the open [project board](https://github.com/panachainy/fiber-101/projects/1)

## Ref

- [Fiber](https://github.com/gofiber/fiber)
- [Live-reload](https://techinscribed.com/5-ways-to-live-reloading-go-applications/)
- [Intro-golang](https://github.com/AnuchitO/intro-golang)
- [Fiber-Golang-YT](https://www.youtube.com/watch?v=Iq2qT0fRhAA&ab_channel=TutorialEdge)
- [Coveralls](https://github.com/mattn/goveralls)
