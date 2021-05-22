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

## Initial project

### Create module

```sh
go mod init demo
```

## Getting started

### Build

`go get`

`go install`

### Run

`go run server.go`

or

`air`

## Ref

> https://github.com/gofiber/fiber

> https://techinscribed.com/5-ways-to-live-reloading-go-applications/
