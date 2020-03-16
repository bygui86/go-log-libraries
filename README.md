
# Go log libraries

Project to compare Golang logging libraries

`WORK IN PROGRESS`

## Libraries (sorted by performances)
1. [zerolog](https://github.com/rs/zerolog)
2. [zap](https://github.com/uber-go/zap)
3. [logrus](https://github.com/sirupsen/logrus)
4. [logur](https://github.com/logur/logur) - `TODO`
5. fmt (package included)
6. log (package included)

## Usage samples
```shell
export GO111MODULE=on
go run zerolog.go
go run zap.go
go run logrus.go
go run logur.go - `TODO`
```

## Benchmark
### Static testing
`TODO`

### Inside an application
#### Run server
```shell
export GO111MODULE=on
go run compare.go
```
#### Run single-print load test
```shell
go get github.com/tsliwowicz/go-wrk
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/zerolog
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/zap
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/logrus
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/logur - `TODO`
```
#### Run multiple-prints load test
```shell
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/zerolog100
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/zap100
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/logrus100
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/logur100 - `TODO`
```
#### Run context switches and syscalls monitor
```shell
sudo strace -c -t -p $(pid)
```

---

## Links

- https://medium.com/techlog/does-logging-cause-cpu-load-a-test-of-all-the-golang-logging-libraries-34052240f90d
- https://medium.com/a-journey-with-go/go-how-zap-package-is-optimized-dbf72ef48f2d
- https://github.com/uber-go/zap
- https://github.com/rs/zerolog
- https://github.com/sirupsen/logrus
- https://github.com/logur/logur
