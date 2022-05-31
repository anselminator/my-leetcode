package main

import "fmt"

func main() {

	s1 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	s2 := []int{10, 15, 20}

	fmt.Println(s1, " , min cost for climbing is:", minCostClimbingStairs(s1))
	fmt.Println(s2, " , min cost for climbing is:", minCostClimbingStairs(s2))
}

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	//var total int = 0
	if l == 2 {
		return min(cost[0], cost[1])
	} else {
		return cost[l-1] + min(minCostClimbingStairs(cost[:l-1]), minCostClimbingStairs(cost[:l-2]))
	}
}
func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
