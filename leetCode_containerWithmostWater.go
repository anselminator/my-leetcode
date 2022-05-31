package main

import "fmt"

// https://leetcode.com/problems/container-with-most-water/
func main() {
	var a = []int{1, 6, 3, 4, 5, 16, 7, 8, 9}
	var b = []int{1, 1}
	var c = []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(a, "has max water ", maxArea(a))
	fmt.Println(b, "has max water ", maxArea(b))
	fmt.Println(c, "has max water ", maxArea(c))
}

func maxArea(height []int) int {
	var bestSoFar int = 0
	var area int
	var n = len(height)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			m := min(height[i], height[j])
			area = m * (j - i)
			if area > bestSoFar {
				bestSoFar = area
			}
		}
	}

	return bestSoFar

}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
