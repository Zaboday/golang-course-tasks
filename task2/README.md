# Результаты бенчмарков

## Реализация на слайсах (пакет slicer)

```
go test -bench=. -benchmem ./pkg/slicer 
goos: linux
goarch: amd64
pkg: main/pkg/slicer
BenchmarkManagerInsert5-4     	33536841	        32.3 ns/op	      48 B/op	       1 allocs/op
BenchmarkManagerInsert50-4    	14205187	        81.6 ns/op	     416 B/op	       1 allocs/op
BenchmarkManagerInsert100-4   	 8306694	       143 ns/op	     896 B/op	       1 allocs/op
BenchmarkManagerDelete5-4     	39361262	        31.8 ns/op	      48 B/op	       1 allocs/op
BenchmarkManagerDelete50-4    	10946396	       110 ns/op	     416 B/op	       1 allocs/op
BenchmarkManagerDelete100-4   	 5010292	       220 ns/op	     896 B/op	       1 allocs/op
PASS
ok  	main/pkg/slicer	8.635s
```

## Двунаправленный связный список (пакет linkedlist)