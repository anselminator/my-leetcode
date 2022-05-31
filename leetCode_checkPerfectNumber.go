package main

import (
	"fmt"
)

func main() {
	fmt.Println(99_999_999, " is a perfect number because  sum(", findDivisors(99_999_999), ")=", 99_999_999)

	for i := 1; i < 20_000; i++ {
		//		fmt.Println("checking ", i)
		if checkPerfectNumber(i) {
			fmt.Println(i, " is a perfect number because  sum(", findDivisors(i), ")=", i)
		}
	}
	fmt.Println(99_999_995, " is a perfect number because  sum(", findDivisors(99_999_995), ")=", 99_999_995)

}

func checkPerfectNumber(num int) bool {
	var sum int = 0
	//var r []int

	for i := 1; i <= (num / 2); i++ {
		if num%i == 0 {
			sum += i
			//r[divs] = i
		}
	}

	return sum == num
}

func findDivisors(num int) []int {
	var r []int
	var divs int = 0
	for i := 1; i <= (num / 2); i++ {
		if num%i == 0 {
			r = append(r, i)
			//r[divs] = i
			divs++
		}
	}
	return r
}
