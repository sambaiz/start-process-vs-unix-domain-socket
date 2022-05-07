# start-process-vs-unix-domain-socket

Experiment of how much the cost of starting a new process can be reduced by passing values by unix domain socket.

```shell
$ make bench
go build -o main .
go test -bench . -benchtime 100x
goos: darwin
goarch: amd64
pkg: github.com/sambaiz/start-process-vs-unix-domain-socket
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkProcess-4                       	     100	   4433338 ns/op
BenchmarkSocketDialEveryTime-4           	     100	  10198513 ns/op
BenchmarkSocketDialOnce-4                	     100	  10109449 ns/op
BenchmarkSocketAlreadyProcessStarted-4   	     100	    101352 ns/op
PASS
ok  	github.com/sambaiz/start-process-vs-unix-domain-socket	7.793s
go test -bench . -benchtime 1000x
goos: darwin
goarch: amd64
pkg: github.com/sambaiz/start-process-vs-unix-domain-socket
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkProcess-4                       	    1000	   5761057 ns/op
BenchmarkSocketDialEveryTime-4           	    1000	   1105874 ns/op
BenchmarkSocketDialOnce-4                	    1000	   1026624 ns/op
BenchmarkSocketAlreadyProcessStarted-4   	    1000	     14375 ns/op
PASS
ok  	github.com/sambaiz/start-process-vs-unix-domain-socket	12.110s
go test -bench . -benchtime 10000x
goos: darwin
goarch: amd64
pkg: github.com/sambaiz/start-process-vs-unix-domain-socket
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
BenchmarkProcess-4                       	   10000	   5112125 ns/op
BenchmarkSocketDialEveryTime-4           	   10000	    205253 ns/op
BenchmarkSocketDialOnce-4                	   10000	    113940 ns/op
BenchmarkSocketAlreadyProcessStarted-4   	   10000	     13312 ns/op
PASS
ok  	github.com/sambaiz/start-process-vs-unix-domain-socket	58.664s
```

## article

- ([ja](https://www.sambaiz.net/article/404/)/[en](https://www.sambaiz.net/en/article/404/)) How faster is sending/receiving values by UNIX domain socket than starting new processes when executing commands
