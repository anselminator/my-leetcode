package main

import (
	"fmt"
	"strconv"
)

func main() {
	r1 := []string{"5", "2", "C", "D", "+"}
	r2 := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	r3 := []string{"1", "C"}
	r4 := []string{"1", "D", "D", "D"}
	//r3 := "UDLR"
	fmt.Println("baseball")
	fmt.Println(r1, calPoints(r1))
	fmt.Println(r2, calPoints(r2))
	fmt.Println(r3, calPoints(r3))
	fmt.Println(r4, calPoints(r4))
}

func calPoints(ops []string) int {
	var scores []int = make([]int, len(ops))
	var idx int = 0
	var total int = 0
	for _, r := range ops {
		switch r {
		case "+":
			scores[idx] = scores[idx-1] + scores[idx-2]
			idx++
		case "D":
			scores[idx] = scores[idx-1] * 2
			idx++
		case "C":
			if idx > 0 {
				idx--
				scores[idx] = 0
			}
		default:
			scores[idx], _ = strconv.Atoi(r)
			idx++
		}
	}
	for j, _ := range scores {
		total += scores[j]
	}
	return total
}
