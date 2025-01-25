package main

import (
	"fmt"
	"time"
)

// Fibonacci
// - F(0) = 0, F(1) = 1
// - For n > 1: F(n) = F(n-1) + F(n-2)

func main() {
	fmt.Printf("--- FIBONACCI ---\n\n")

	n := 40
	fmt.Printf("n: %v\n", n)

	// iterative fibonacci call
	startTime := time.Now()
	result := fibonacciIterative(n)
	duration := time.Since(startTime)

	fmt.Printf("\niterative fibonacci result: %v\n", result)
	fmt.Printf("execution time: %v\n", duration)

	// recursive fibonacci call
	startTime = time.Now()
	result = fibonacciRecursive(n)
	duration = time.Since(startTime)

	fmt.Printf("\nrecursive fibonacci result: %v\n", result)
	fmt.Printf("execution time: %v\n", duration)

	// recursive fibonacci with memoization call
	startTime = time.Now()
	result = fibonacciMemoization(n)
	duration = time.Since(startTime)

	fmt.Printf("\nrecursive fibonacci with memoization result: %v\n", result)
	fmt.Printf("cache: %v\n", memo)
	fmt.Printf("execution time: %v\n", duration)
}

func fibonacciIterative(n int) int {
	// base case
	if n <= 1 {
		return n
	}

	// iterative case
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func fibonacciRecursive(n int) int {
	// base case
	if n <= 1 {
		return n
	}

	// iterative case
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

var memo = make(map[int]int)

func fibonacciMemoization(n int) int {
	// base case
	if n <= 1 {
		return n
	}

	// iterative case
	if _, exists := memo[n]; !exists {
		memo[n] = fibonacciMemoization(n-1) + fibonacciMemoization(n-2)
	}
	return memo[n]
}
