
# go-log-libraries

## Libraries (sorted by performances)
1. [zerolog](https://github.com/rs/zerolog)
2. [zap](https://github.com/uber-go/zap)
3. [logrus](https://github.com/sirupsen/logrus)
4. fmt (package included)
5. log (package included)

## Single samples
```
export GO111MODULE=on
go run zerolog.go
go run zap.go
go run logrus.go
```

## Comparison
### Run server
```
export GO111MODULE=on
go run compare.go
```
### Run single-print load test
```
go get github.com/tsliwowicz/go-wrk
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/zerolog
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/zap
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/logrus
```
### Run multiple-prints load test
```
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/zerolog100
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/zap100
go-wrk -c 1000 -d 10 -T 20000 http://0.0.0.0:8080/logrus100
```
### Run context switches and syscalls monitor
```
sudo strace -c -t -p $(pid)
```

---

## Links

* https://medium.com/techlog/does-logging-cause-cpu-load-a-test-of-all-the-golang-logging-libraries-34052240f90d
* https://medium.com/a-journey-with-go/go-how-zap-package-is-optimized-dbf72ef48f2d
