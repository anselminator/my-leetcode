package main

import "fmt"

func main() {
	for i := 0; i < 254; i++ {
		fmt.Println(i, " is a power of three: ", isPowerOfThree(i))
	}
}
func isPowerOfThree(n int) bool {
	for n >= 3 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	if n == 1 {
		return true
	} else {
		return false
	}
}
