package main

import "fmt"

type mapval struct {
	val, idx int
}

func main() {
	a := []int{2, 7, 11, 15}
	b := []int{3, 2, 4}
	c := []int{3, 3}
	ta := 9
	tb := 6
	tc := 6

	fmt.Println("a=", a, " , target", ta, " result=", twoSum(a, ta))
	fmt.Println("b=", b, " , target", tb, " result=", twoSum(b, tb))
	fmt.Println("c=", c, " , target", tc, " result=", twoSum(c, tc))
}

func twoSum(nums []int, target int) []int {
	var dic map[int]mapval = make(map[int]mapval)
	//	var l int = len(nums)
	//	var i, j int = 0, 0
	//	dic[1] = mapval{2, 3}
	for i, val := range nums {
		a, ok := dic[val]
		if ok {
			return []int{i, a.idx}
		} else {
			dic[target-val] = mapval{target - val, i}
		}
	}
	return nil
}
func numOfPairs(nums []string, target string) int {
	var l int = len(nums)
	var i, j, count int = 0, 0, 0
	for i = 0; i < l; i++ {
		for j = 0; j < l; j++ {
			if nums[i]+nums[j] == target {
				count++
			}
		}
	}
	return count
}
