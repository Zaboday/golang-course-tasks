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


```
go test -bench=. -benchmem ./pkg/linkedlist
goos: linux
goarch: amd64
pkg: main/pkg/linkedlist
BenchmarkLinkedListInsert5-4     	 5304624	       220 ns/op	     192 B/op	       6 allocs/op
BenchmarkLinkedListInsert50-4    	  342417	      3338 ns/op	    1632 B/op	      51 allocs/op
BenchmarkLinkedListInsert100-4   	  117426	     10002 ns/op	    3232 B/op	     101 allocs/op
BenchmarkLinkedListDelete5-4     	 6326252	       186 ns/op	     160 B/op	       5 allocs/op
BenchmarkLinkedListDelete50-4    	  339898	      3270 ns/op	    1600 B/op	      50 allocs/op
BenchmarkLinkedListDelete100-4   	  117088	      9977 ns/op	    3200 B/op	     100 allocs/op
PASS
ok  	main/pkg/linkedlist	7.651s
```