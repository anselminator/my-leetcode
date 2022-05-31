package main

import "fmt"

func main() {
	s1 := "abcde"
	g1 := "cdeab"
	s2 := "abcde"
	g2 := "abced"
	fmt.Println(s1, g1, " is ", rotateString(s1, g1))
	fmt.Println(s2, g2, " is ", rotateString(s2, g2))
}

func rotateString(s string, goal string) bool {
	var l int = len(s)
	for i := 0; i < l; i++ {
		r := string(s[0])
		s = s[1:l] + r
		fmt.Println(r, s, goal)
		if goal == s {
			return true
		}
	}
	return false
}
