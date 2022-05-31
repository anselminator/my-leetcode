package main

import "fmt"

func main() {
	a1 := 4
	a2 := 10
	a3 := 20
	a4 := 100
	fmt.Println("Leetcode Bank.....")
	fmt.Println("atfter ", a1, " days, ", totalMoney(a1), " in leetcodes bank")
	fmt.Println("atfter ", a2, " days, ", totalMoney(a2), " in leetcodes bank")
	fmt.Println("atfter ", a3, " days, ", totalMoney(a3), " in leetcodes bank")
	fmt.Println("atfter ", a4, " days, ", totalMoney(a4), " in leetcodes bank")
}
func totalMoney(n int) int {
	var total int = 0
	var weeks int = 0
	var days int = 1
	for i := 1; i <= n; i++ {
		total += (weeks + days)
		//		fmt.Println("week +", weeks, " days ", days, "total =", total)
		days++
		if i%7 == 0 {
			//total += weeks + days
			weeks++
			days = 1
			//			fmt.Println("week +", weeks, " days ", days, " reset. Total=", total)
		}
	}
	//	fmt.Println("week +", weeks, " days ", days, " final")

	return total
}
