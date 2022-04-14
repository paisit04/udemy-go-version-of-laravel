# Let's Build a Go version of Laravel

Ref: https://www.udemy.com/course/lets-build-a-go-version-of-laravel

## testing
```sh
$ go test .
$ go test -cover .
```

## integration testing
```sh
$ go test . --tags integration --count=1
$ go test -cover . --tags integration
$ go test -coverprofile=coverage.out . --tags integration
$ go tool cover -html=coverage.out
```

## myapp

```sh
$ go get -u github.com/paisit04/celeritas
$ go mod vendor
$ make run
```
