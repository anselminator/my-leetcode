package main

import "fmt"

func main() {
	for i := 0; i < 65; i++ {

		fmt.Println(i, " is perfect square : ", isPerfectSquare(i))
	}
}

func isPerfectSquare(num int) bool {
	for i := 1; i <= num>>1; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
