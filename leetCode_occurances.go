package main

import "fmt"

func main() {

	a := []int{3, 2, 3}
	b := []int{1}
	c := []int{1, 2}
	fmt.Println(a, "=", majorityElement(a))
	fmt.Println(b, "=", majorityElement(b))
	fmt.Println(c, "=", majorityElement(c))
}
func majorityElement(nums []int) []int {
	var l int = len(nums)
	var limit int = 0
	//var occ []int = make([]int,l)
	occ := map[int]int{}
	var res []int

	limit = (l / 3)

	/*	if l%3 == 0 {
			limit = (l / 3)
		} else {
			limit = (l / 3) + 1

		}
	*/
	for _, v := range nums {
		occ[v] = occ[v] + 1
	}
	fmt.Println(occ, nums, limit)

	for k, v := range occ {
		//	r, ok := occ[v]
		if v > limit {
			res = append(res, k)
		}

	}
	return res
}
