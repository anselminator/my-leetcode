package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("fizzbuzz")
	e := fizzBuzz(2000)
	fmt.Println(e)
}

func fizzBuzz(n int) []string {
	answer := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer[i-1] = "FizzBuzz"
		} else if i%3 == 0 || i%5 == 0 {
			if i%3 == 0 {
				answer[i-1] = "Fizz"
			} else {
				answer[i-1] = "Buzz"
			}
		} else {
			answer[i-1] = strconv.Itoa(i)
		}
	}
	return answer
}
