# Результаты бенчмарков

## Реализация на слайсах (пакет slicer)

```
go test -bench=. -benchmem ./pkg/slicer
goos: linux
goarch: amd64
pkg: main/pkg/slicer
BenchmarkManagerInsert5-4     	28835568	        37.8 ns/op	      48 B/op	       1 allocs/op
BenchmarkManagerInsert50-4    	 8576161	       135 ns/op	     416 B/op	       1 allocs/op
BenchmarkManagerInsert100-4   	 4956202	       237 ns/op	     896 B/op	       1 allocs/op
BenchmarkManagerDelete5-4     	35437718	        32.1 ns/op	      48 B/op	       1 allocs/op
BenchmarkManagerDelete50-4    	10511185	       111 ns/op	     416 B/op	       1 allocs/op
BenchmarkManagerDelete100-4   	 5325783	       209 ns/op	     896 B/op	       1 allocs/op
PASS
ok  	main/pkg/slicer	7.650s
```

## Двунаправленный связный список (пакет linkedlist)