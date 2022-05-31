package main

import "fmt"

func main() {

	a := []int{3, 2, 3}
	b := []int{2, 2, 1, 1, 1, 2, 2}
	c := []int{1, 2}
	fmt.Println(a, "=", majorityElement(a))
	fmt.Println(b, "=", majorityElement(b))
	fmt.Println(c, "=", majorityElement(c))
}
func majorityElement(nums []int) int {
	var l int = len(nums)
	var limit int = 0
	occ := map[int]int{}
	var max int = 0
	var res int = 0

	limit = (l / 2)

	for _, v := range nums {
		occ[v] = occ[v] + 1
	}

	for key, val := range occ {
		if val > limit && val > max {
			max = val
			res = key
		}

	}
	return res
}
