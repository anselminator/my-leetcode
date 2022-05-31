package main

import "fmt"

func main() {
	for i := 1; i < 130; i++ {
		fmt.Printf("%0b ", i)
		fmt.Println("numOfSTeps(", i, " )= ", numberOfSteps(i))
	}
}

func numberOfSteps(num int) int {
	var count int = 0
	for num > 0 {
		if num%2 == 0 {
			count++
			num = num >> 1
		} else {
			count++
			num = num - 1
		}
	}
	return count
}
