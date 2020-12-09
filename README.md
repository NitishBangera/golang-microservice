## Build Steps

1. Install dependencies : `go mod tidy && go mod vendor`
3. Change the kafka addresses in main.go
4. Run the code : `go run main.go`

## Docker Build
```sh
$ docker build -t test .
$ docker run -v /etc/config.ini:/etc/config.ini -i -t test
```