package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{6, 2, 7, 3}
	af := 1
	bf := 4
	fmt.Println(a, " ", af, " gives ", decode(a, af))
	fmt.Println(b, " ", bf, " gives ", decode(b, bf))
}
func decode(encoded []int, first int) []int {
	var res []int = make([]int, len(encoded)+1)
	res[0] = first
	for i := 0; i < len(encoded); i++ {
		res[i+1] = res[i] ^ encoded[i]
	}
	return res
}
