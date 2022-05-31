package main

import "fmt"

type oc struct {
	times int
	index int
}

func main() {
	s := "leetcode"
	s2 := "loveleetcode"
	s3 := "aabb"

	fmt.Println(s, firstUniqChar(s))
	fmt.Println(s2, firstUniqChar(s2))
	fmt.Println(s3, firstUniqChar(s3))
}
func firstUniqChar(s string) int {
	var occur map[rune]oc = make(map[rune]oc)
	var res = 99999
	for i, r := range s {
		v, ok := occur[r]
		if ok {
			occur[r] = oc{v.times + 1, v.index}
		} else {
			occur[r] = oc{1, i}
		}
	}
	for _, vl := range occur {
		if vl.times == 1 && vl.index < res {
			res = vl.index
		}
	}
	if res < 99999 {
		return res
	} else {
		return -1
	}

}
