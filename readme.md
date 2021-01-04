# Benchmark dynamic programming en Go

Toy project to practice benchmarking with the Go programming language.

Compare memory and runtime of three [`fibonacci`](https://en.wikipedia.org/wiki/Fibonacci_number) functions.

1. The naive recursive approach
2. An iterative approach, using a slice.
3. An optimized iterative approach with a fixed number of variables.

### Benchmark 

Run the benchmark with:

````shell
go test -bench=. -benchmem > benchmark/raw.txt
````

Results on `goos: darwin` and `goarch: amd64` available in `./benchmark/raw.txt`.

### Useful links

- [How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go), from Dave Cheney
- [Cracking the Coding Interview](https://www.crackingthecodinginterview.com/), chapter 8, Dynamic Programming


Benchmark analysis in [my article](https://fpaupier.dev/blog/fibonacci-benchmark/) (french)
