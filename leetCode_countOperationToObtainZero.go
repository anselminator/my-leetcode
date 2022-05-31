package main

import "fmt"

func main() {
	fmt.Println("numOfSTepstoZero(2,3)= ", countOperations(2, 3))
	fmt.Println("numOfSTepstoZero(10,10)= ", countOperations(10, 10))
}

func countOperations(num1 int, num2 int) int {
	var count int = 0
	for num1 > 0 && num2 > 0 {
		if num1 >= num2 {
			count++
			num1 = num1 - num2
		} else {
			count++
			num2 = num2 - num1
		}
	}
	return count
}
