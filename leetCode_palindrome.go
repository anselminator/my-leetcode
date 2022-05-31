package main

import (
	"fmt"
	"strconv"
)

func main() {
	var bla int = 123454321
	var bla2 int = 91768576
	var bla3 int = 87967876978
	var bla4 int = 879678876978
	var bla5 int = 896782876978

	var s1 string = "lecco mio"

	fmt.Println("leck mich", strconv.Itoa(bla))
	fmt.Println(reverseStr("leck mich"), strconv.Itoa(42))
	fmt.Println(reverseStr(s1), strconv.Itoa(-420))
	s2 := reverseStr(s1)
	fmt.Println("jetzt aber:", string(s2))
	s3 := fmt.Sprintf("blalal %i bbbbbb %d cndjcnd %s", 4, 5, "fuck")
	fmt.Println(s3)
	fmt.Println(bla, " is palindrome :", isPalindrome(bla))
	fmt.Println(bla2, " is palindrome :", isPalindrome(bla2))
	fmt.Println(bla3, " is palindrome :", isPalindrome(bla3))
	fmt.Println(bla4, " is palindrome :", isPalindrome(bla4))
	fmt.Println(bla5, " is palindrome :", isPalindrome(bla5))
	fmt.Println("0", " is palindrome :", isPalindrome(0))
	fmt.Println("123", " is palindrome :", isPalindrome(112232211))
	fmt.Println("124", " is palindrome :", isPalindrome(12454321))
}

/*func isPalindrome(x int) bool {
	return strconv.Itoa(x) == reverseStr(strconv.Itoa(x))
}*/
func reverseStr(s string) (r string) {
	for _, c := range s {
		r = string(c) + r
	}
	return r
}
