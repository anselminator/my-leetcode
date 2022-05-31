package main

import "fmt"

func main() {
	s1 := "Hello World"
	s2 := "   fly me   to   the moon  "
	s3 := "luffy is still joyboy"
	s4 := "aa"
	s5 := "1234"
	fmt.Println(len(s1), s1, lengthOfLastWord(s1))
	fmt.Println(len(s2), s2, lengthOfLastWord(s2))
	fmt.Println(len(s3), s3, lengthOfLastWord(s3))
	fmt.Println(len(s4), s4, lengthOfLastWord(s4))
	fmt.Println(len(s5), s5, lengthOfLastWord(s5))
}

func lengthOfLastWord(s string) int {
	var lastspace int = 0
	var trim int = 0
	//	var lastlength int = 0
	// trim whitespace off the end
	for trim = len(s); rune(s[trim-1]) == ' '; trim-- {
	}
	fmt.Println(len(s)-(trim), "whitespaces snipped off. trim =", trim)

	for idx, r := range s {
		if r == ' ' && idx < trim {
			lastspace = idx + 1
		}
	}

	return trim - lastspace
}
