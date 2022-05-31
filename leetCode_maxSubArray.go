package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, -20, 6, 7, 8, 9}
	b := []int{5, 4, -1, 7, 8}
	c := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	d := []int{1}
	fmt.Println("bleh")
	fmt.Println("maxSubarray", a, maxSubArray(a))
	fmt.Println("maxSubarray", b, maxSubArray(b))
	fmt.Println("maxSubarray", c, maxSubArray(c))
	fmt.Println("maxSubarray", d, maxSubArray(d))
	fmt.Println("subsum 0->3 = ", subSum(a, 0, 3))
	fmt.Println("subsum 5->8 = ", subSum(a, 5, 8))
	fmt.Println("subsum 4->4 = ", subSum(a, 4, 4))
	fmt.Println("subsum 0->4 = ", subSum(a, 0, 4))
}
func maxSubArray(nums []int) int {
	// check out Kadane for a much better version
	//	var bestStart int = 0
	//	var bestEnd int = 0
	var bestSum int = 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sub := subSum(nums, i, j)
			if sub > bestSum {
				bestSum = sub
				//				bestStart = i
				//				bestEnd = j
			}
		}
	}
	return bestSum
}

func subSum(nums []int, from int, to int) int {
	var s int = 0
	for i := from; i <= to; i++ {
		s += nums[i]
	}
	return s
}
