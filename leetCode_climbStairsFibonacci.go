package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	for i := 1; i < 45; i++ {
		fmt.Println(i, " steps, ", climbStairs(i), " different ways to climb staircase")
	}
	t3 := time.Since(t1)

	t2 := time.Now()
	for i := 1; i < 45; i++ {
		fmt.Println(i, " steps, ", climbStairs2(i), " different ways to climb staircase")
	}
	t4 := time.Since(t2)

	fmt.Println("recursive: ", t3, ", iterative: ", t4)
}
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return climbStairs(n-2) + climbStairs(n-1)
}

func climbStairs2(n int) int {
	//var res int = 0
	var prev1, prev2, tmp int
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	prev2 = 2
	prev1 = 1
	for i := 3; i <= n; i++ {
		tmp = prev1 + prev2
		prev1 = prev2
		prev2 = tmp
	}
	return prev2
}

func tribonacci(n int) int {
	//var res int = 0
	var prev1, prev2, prev3, tmp int
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	prev3 = 1
	prev2 = 1
	prev1 = 0
	for i := 3; i <= n; i++ {
		tmp = prev1 + prev2 + prev3
		prev1 = prev2
		prev2 = prev3
		prev3 = tmp
	}
	return prev3
}
func fib(n int) int {
	var prev1, prev2, tmp int
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	prev2 = 1
	prev1 = 0
	for i := 2; i <= n; i++ {
		tmp = prev1 + prev2
		prev1 = prev2
		prev2 = tmp
	}
	return prev2
}
