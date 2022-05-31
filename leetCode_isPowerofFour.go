package main

import "fmt"

func main() {
	for i := 1; i <= 300; i++ {

		fmt.Printf("%d ---- %b", i, i)
		fmt.Println(" is a power of 4 ", isPowerOfFour(i))
	}
}
func isPowerOfFour(n int) bool {
	var res bool = true
	if n <= 0 {
		return false
	}
	for n >= 4 {
		res = res && (n%2 == 0)
		n = n >> 2
		res = res && (n%2 == 0)
	}
	return res
}
