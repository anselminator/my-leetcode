package main

import "fmt"

func main() {
	a := []int{5, 7, 7, 8, 8, 10}
	b := []int{}
	c := []int{1}
	d := []int{2, 2}
	e := []int{4, 4, 4, 4, 4, 4}

	fmt.Println(a)
	fmt.Println("searcRange ", 8, searchRange(a, 8))
	fmt.Println("searcRange ", 6, searchRange(b, 0))
	fmt.Println("searcRange ", 6, searchRange(c, 1))
	fmt.Println("searcRange ", 6, searchRange(d, 2))
	fmt.Println("searcRange ", 6, searchRange(e, 4))
}

func searchRange(nums []int, target int) []int {
	var l int = len(nums)
	var ll int = 0
	var rr int = l - 1
	var m int = 0
	var il, iu int = 0, 0
	if l == 0 {
		return []int{-1, -1}
	}
	if l == 1 && target == nums[0] {
		return []int{0, 0}
	}
	if target < nums[0] {
		return []int{-1, -1}
	}
	if target > nums[l-1] {
		return []int{-1, -1}
	}
	for ll <= rr {
		//get middle
		m = (ll + rr) >> 1

		if nums[m] > target {
			rr = m - 1
			continue
		} else if nums[m] < target {
			ll = m + 1
			continue
		} else if nums[m] == target { //we found it
			il = m
			iu = m
			for m > 0 && nums[m] == target {
				m--
			}
			if nums[m] == target {
				il = m
			} else {
				il = m + 1
			}
			m = iu
			for m < l && nums[m] == target {
				m++
			}
			if m == l {
				m--
			}
			if nums[m] == target {
				iu = m
			} else {
				iu = m - 1
			}
			return []int{il, iu}
		}
	}
	return []int{-1, -1}
}
