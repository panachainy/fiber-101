# Fiber-101

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

### Create module

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

`go get`

`go install`

`go mod download`

#### Clean dependency

`go mod tidy -v`

### Run

`go run server.go`

or

`air`

## Docker

### Docker build

```sh
docker build -t fiber .
```

### Docker test run

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

## Ref

> https://github.com/gofiber/fiber

> https://techinscribed.com/5-ways-to-live-reloading-go-applications/

> https://github.com/AnuchitO/intro-golang

> https://www.youtube.com/watch?v=Iq2qT0fRhAA&ab_channel=TutorialEdge

> https://github.com/mattn/goveralls
