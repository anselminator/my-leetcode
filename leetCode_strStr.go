package main

import "fmt"

func main() {
	s1 := "helllo world"
	fmt.Println(s1)
	fmt.Println(s1[1:4])
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("abc", "c"))
	fmt.Println(strStr("aaa", "a"))

}
func strStr(haystack string, needle string) int {
	var ndl int = len(needle)
	var hsl int = len(haystack)
	var pos int = -1
	if needle == haystack {
		return 0
	}
	for i := 0; i <= hsl-ndl; i++ {
		if needle == haystack[i:i+ndl] {
			pos = i
			break
		}

	}
	return pos
}
