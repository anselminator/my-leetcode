package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{4, 9, 9, 9}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
	fmt.Println(plusOne([]int{9, 8, 9, 9}))
	fmt.Println(plusOne([]int{9, 7, 9, 9}))
	fmt.Println(plusOne([]int{9, 9, 9, 1}))
}

func plusOne(digits []int) []int {
	l := len(digits)
	if digits[l-1] < 9 {
		digits[l-1]++
		return digits
	} else {
		if l == 1 {
			return []int{1, 0}
		} else {
			digits = append(plusOne(digits[:l-1]), 0)
		}
	}
	return digits
}
