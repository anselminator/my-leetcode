package main

import "fmt"

func main() {
	for i := 0; i < 4200; i++ {
		fmt.Println("sqrt(", i, ")=", mySqrt(i))
	}
	fmt.Println("sqrt(100_000)=", mySqrt(100_000))
	fmt.Println("sqrt(999_999_999)=", mySqrt(999_999_999))
}

func mySqrt(x int) int {
	switch x {
	case 0:
		return 0
	case 1, 2:
		return 1
	default:
		for i := 0; i < x; i++ {
			if i*i >= x {
				if i*i == x {
					return i
				} else {
					return i - 1
				}
			}
		}
	}
	return -1
}
