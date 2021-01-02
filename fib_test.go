package main

import "testing"

// For all test store memory consumed and time elapsed

// Test Fib for various values of n
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		fib(n)
	}
}
func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib4(b *testing.B)  { benchmarkFib(b, 4) }
func BenchmarkFib5(b *testing.B)  { benchmarkFib(b, 5) }
func BenchmarkFib6(b *testing.B)  { benchmarkFib(b, 6) }
func BenchmarkFib7(b *testing.B)  { benchmarkFib(b, 7) }
func BenchmarkFib8(b *testing.B)  { benchmarkFib(b, 8) }
func BenchmarkFib9(b *testing.B)  { benchmarkFib(b, 9) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib11(b *testing.B) { benchmarkFib(b, 11) }
func BenchmarkFib12(b *testing.B) { benchmarkFib(b, 12) }
func BenchmarkFib13(b *testing.B) { benchmarkFib(b, 13) }
func BenchmarkFib14(b *testing.B) { benchmarkFib(b, 14) }
func BenchmarkFib15(b *testing.B) { benchmarkFib(b, 15) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib25(b *testing.B) { benchmarkFib(b, 25) }
func BenchmarkFib30(b *testing.B) { benchmarkFib(b, 30) }
func BenchmarkFib35(b *testing.B) { benchmarkFib(b, 35) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }
func BenchmarkFib41(b *testing.B) { benchmarkFib(b, 41) }

// Test FibIter for various values of n
func benchmarkFibIter(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		fibIter(n)
	}
}
func BenchmarkFibIter1(b *testing.B)  { benchmarkFibIter(b, 1) }
func BenchmarkFibIter2(b *testing.B)  { benchmarkFibIter(b, 2) }
func BenchmarkFibIter3(b *testing.B)  { benchmarkFibIter(b, 3) }
func BenchmarkFibIter4(b *testing.B)  { benchmarkFibIter(b, 4) }
func BenchmarkFibIter5(b *testing.B)  { benchmarkFibIter(b, 5) }
func BenchmarkFibIter6(b *testing.B)  { benchmarkFibIter(b, 6) }
func BenchmarkFibIter7(b *testing.B)  { benchmarkFibIter(b, 7) }
func BenchmarkFibIter8(b *testing.B)  { benchmarkFibIter(b, 8) }
func BenchmarkFibIter9(b *testing.B)  { benchmarkFibIter(b, 9) }
func BenchmarkFibIter10(b *testing.B) { benchmarkFibIter(b, 10) }
func BenchmarkFibIter11(b *testing.B) { benchmarkFibIter(b, 11) }
func BenchmarkFibIter12(b *testing.B) { benchmarkFibIter(b, 12) }
func BenchmarkFibIter13(b *testing.B) { benchmarkFibIter(b, 13) }
func BenchmarkFibIter14(b *testing.B) { benchmarkFibIter(b, 14) }
func BenchmarkFibIter15(b *testing.B) { benchmarkFibIter(b, 15) }
func BenchmarkFibIter20(b *testing.B) { benchmarkFibIter(b, 20) }
func BenchmarkFibIter25(b *testing.B) { benchmarkFibIter(b, 25) }
func BenchmarkFibIter30(b *testing.B) { benchmarkFibIter(b, 30) }
func BenchmarkFibIter35(b *testing.B) { benchmarkFibIter(b, 35) }
func BenchmarkFibIter40(b *testing.B) { benchmarkFibIter(b, 40) }
func BenchmarkFibIter41(b *testing.B) { benchmarkFibIter(b, 41) }

// Test FibOpti for various values of n
func benchmarkFibOpti(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		fibOpti(n)
	}
}
func BenchmarkFibOpti1(b *testing.B)  { benchmarkFibOpti(b, 1) }
func BenchmarkFibOpti2(b *testing.B)  { benchmarkFibOpti(b, 2) }
func BenchmarkFibOpti3(b *testing.B)  { benchmarkFibOpti(b, 3) }
func BenchmarkFibOpti4(b *testing.B)  { benchmarkFibOpti(b, 4) }
func BenchmarkFibOpti5(b *testing.B)  { benchmarkFibOpti(b, 5) }
func BenchmarkFibOpti6(b *testing.B)  { benchmarkFibOpti(b, 6) }
func BenchmarkFibOpti7(b *testing.B)  { benchmarkFibOpti(b, 7) }
func BenchmarkFibOpti8(b *testing.B)  { benchmarkFibOpti(b, 8) }
func BenchmarkFibOpti9(b *testing.B)  { benchmarkFibOpti(b, 9) }
func BenchmarkFibOpti10(b *testing.B) { benchmarkFibOpti(b, 10) }
func BenchmarkFibOpti11(b *testing.B) { benchmarkFibOpti(b, 11) }
func BenchmarkFibOpti12(b *testing.B) { benchmarkFibOpti(b, 12) }
func BenchmarkFibOpti13(b *testing.B) { benchmarkFibOpti(b, 13) }
func BenchmarkFibOpti14(b *testing.B) { benchmarkFibOpti(b, 14) }
func BenchmarkFibOpti15(b *testing.B) { benchmarkFibOpti(b, 15) }
func BenchmarkFibOpti20(b *testing.B) { benchmarkFibOpti(b, 20) }
func BenchmarkFibOpti25(b *testing.B) { benchmarkFibOpti(b, 25) }
func BenchmarkFibOpti30(b *testing.B) { benchmarkFibOpti(b, 30) }
func BenchmarkFibOpti35(b *testing.B) { benchmarkFibOpti(b, 35) }
func BenchmarkFibOpti40(b *testing.B) { benchmarkFibOpti(b, 40) }
func BenchmarkFibOpti41(b *testing.B) { benchmarkFibOpti(b, 41) }
