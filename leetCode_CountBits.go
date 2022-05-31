package main

import "fmt"

func main() {
	fmt.Println("1 bits in ", 2, " = ", countBits(2))
	fmt.Println("1 bits in ", 5, " = ", countBits(5))
	fmt.Println("1 bits in ", 10, " = ", countBits(10))
	fmt.Println("1 bits in ", 16, " = ", countBits(16))
	fmt.Println("1 bits in ", 256, " = ", countBits(256))
}

func countBits(n int) []int {
	var res []int = make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = cntBits(i)
	}
	return res
}

func cntBits(l int) int {
	var res int = 0
	for l != 0 {
		if l%2 == 1 {
			res++
		}
		l = l / 2
	}
	return res
}
