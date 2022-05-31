package main

import (
	"fmt"
)

func main() {
	//var bla2 int = 91768576
	//var s1 string = "lecco mio"
	b := []int{1, 2, 2, 6, 7, 2, 9, 10, 2, 2, 1, 12, 3, 4, 6}
	a := []int{1, 0, 1, 3, 0, 7, 8, 9, 5, 5, 5, 5, 5, 5}
	c := []int{3, 2, 2, 3}
	d := []int{0, 1, 2, 2, 3, 0, 4, 2}
	e := []int{2}
	fmt.Println("b", b, len(b))
	fmt.Println("a", a, len(a))
	fmt.Println("c", c, len(c))
	fmt.Println("d", d, len(d))
	fmt.Println("e", e, len(e))
	s1 := removeElement(b, 50)
	s2 := removeElement(a, 89)
	s3 := removeElement(c, 3)
	s4 := removeElement(d, 2)
	s5 := removeElement(e, 3)
	//	removeElement({2,3,4,5,7,8}, 2)
	fmt.Println("b", b, s1)
	fmt.Println("a", a, s2)
	fmt.Println("c", c, s3)
	fmt.Println("d", d, s4)
	fmt.Println("e", e, s5)

}

func removeElement(nums []int, val int) int {
	var found int = 0
	var temp int = 0
	var newLength = 0

	for i, v := range nums {
		if v == val { // mark element for removal
			found++
			nums[i] = -1
		}
	}
	newLength = len(nums) - (found)

	for i := 0; i < newLength; i++ {
		if nums[i] == -1 { // if a "hole" is found, scan for next non-removed element
			j := i
			for nums[j] == -1 {
				j++
			}
			// then swap em
			temp = nums[i]
			nums[i] = nums[j]
			nums[j] = temp
		}
	}
	return newLength
}
