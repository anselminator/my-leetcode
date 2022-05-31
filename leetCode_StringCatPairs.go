package main

import "fmt"

type mapval struct {
	val, idx int
}

func main() {
	a := []string{"777", "7", "77", "77"}
	b := []string{"123", "4", "12", "34"}
	c := []string{"1", "1", "1"}
	ta := "7777"
	tb := "1234"
	tc := "11"

	fmt.Println("a=", a, " , target", ta, " result=", numOfPairs(a, ta))
	fmt.Println("b=", b, " , target", tb, " result=", numOfPairs(b, tb))
	fmt.Println("c=", c, " , target", tc, " result=", numOfPairs(c, tc))
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
			if i != j && nums[i]+nums[j] == target {
				count++
			}
		}
	}
	return count
}
