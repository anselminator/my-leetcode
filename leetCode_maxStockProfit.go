package main

import "fmt"

func main() {
	e1 := []int{7, 1, 5, 3, 6, 4}
	e2 := []int{7, 6, 4, 3, 1}
	fmt.Println(e1, " max proft is: ", maxProfit(e1))
	fmt.Println(e1, " max proft is: ", maxProfit(e2))
}
func maxProfit(prices []int) int {
	var max, min, maxi, mini int
	var priceMap map[int]int = make(map[int]int) // key=price, value=day

	for i, v := range prices {
		priceMap[v] = i
		if prices[i] > max {
			max = prices[i]
			maxi = i
		}
		if prices[i] < min {
			min = prices[i]
			mini = i
		}
	}
	fmt.Println("max", max, "@", maxi, "; min", min, "@", mini)
	//	if maxi
	for price, day := range priceMap {

	}

}
