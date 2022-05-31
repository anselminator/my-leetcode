package main

import (
	"fmt"
)

func main() {
	s1 := "abcd"
	g1 := "abcde"
	s2 := ""
	g2 := "y"
	fmt.Println(s1, " ", g1, " = ", string(findTheDifference(s1, g1)))
	fmt.Println(s2, " ", g2, " = ", string(findTheDifference(s2, g2)))

}
func findTheDifference(s string, t string) byte {
	for _, r := range s {
		t, _ = find(r, t)
		//		fmt.Println(i, r, t)

	}
	//	fmt.Println(t)
	return byte(rune(t[0]))
}

func find(r rune, s string) (string, bool) {
	for i, a := range s {
		if a == r {
			r := s[:i]
			r2 := s[i+1:]
			return r + r2, true
		}
	}
	return s, false
}
