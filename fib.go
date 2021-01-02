package main

import (
	"fmt"
)

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func fibIter(n int) int {
	dp := make([]int, n+1, n+1)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fibOpti(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	pprev := 0
	prev := 1
	var cur int
	for i := 2; i < n+1; i++ {
		cur = prev + pprev
		pprev = prev
		prev = cur
	}
	return cur
}

func main() {
	fmt.Println(fib(4))
	fmt.Println(fibIter(4))
	fmt.Println(fibOpti(4))
}
