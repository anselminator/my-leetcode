package main

import (
	"fmt"
)

func main() {
	strs1 := []string{"flower", "flow", "floight"}
	strs2 := []string{"dog", "racecar", "car"}

	fmt.Println(strs1)
	fmt.Println(strs2)
	a := longestCommonPrefix(strs1)
	b := longestCommonPrefix(strs2)
	fmt.Println("strs1 longest =", a)
	fmt.Println("strs2 longest =", b)
	fmt.Println(b)

}

func longestCommonPrefix(strs []string) string {
	var mindex int = 0
	var minLen int = 300
	var last = 0
	var same bool = true

	n := len(strs)
	for idx, v := range strs {
		if len(v) < minLen {
			minLen = len(v)
			mindex = idx
		}
	}
	for idx, r := range strs[mindex] {
		fmt.Println(idx, r)
		for _, words := range strs {
			same = same && r == rune(words[idx])
		}
		if !same {
			break
		} else {
			last++
		}
	}
	fmt.Println(mindex, n)
	return strs[mindex][:last]
}
