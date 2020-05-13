<!-- .slide: class="with-code-bg-dark" -->

# pprof

## Outil de profiling built-in

```shell
$ go test -v -bench=. -memprofile=prof.mem
BenchmarkTaskControllerGet-8 200000 5988 ns/op
BenchmarkTaskControllerPost-8 100000 15374 ns/op
BenchmarkHugeMemoryAllocation-8 30000 55522 ns/op
PASS

$ go tool pprof --alloc_space prof.mem
Entering interactive mode (type "help" for commands)
(pprof) top 3
13.81GB of 14.27GB total (96.77%)
Dropped 38 nodes (cum <= 0.07GB)
Showing top 3 nodes out of 29 (cum >= 0.07GB)
flat flat% sum% cum cum%
13.41GB 93.98% 93.98% 13.41GB 93.98% github.com/Sfeir/golang-200/web.BenchmarkHugeMemoryAllocation
0.12GB 0.87% 94.85% 0.12GB 0.87% net/http/httptest.cloneHeader
0.11GB 0.77% 95.62% 0.11GB 0.77% net/textproto.MIMEHeader.Set
```

Notes:
OFU

go 1.9 pas besoin du binaire que le prof.mem

plein d’outils pour diagnostiquer les problèmes

comme jconsole

PRODUCTION READY
