package main

import "fmt"

var phone []string = []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func main() {
	//	phone:=["","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"]
	a := "23"
	b := ""
	c := "2"
	fmt.Println(phone)
	fmt.Println("res", a, letterCombinations(a))
	fmt.Println("res", b, letterCombinations(b))
	fmt.Println("res", c, letterCombinations(c))
}

func letterCombinations(digits string) []string {
	var res []string = make([]string, len(digits))
	for k, v := range digits {
		fmt.Println(k, v, string(v), phone[int(v-'1')])
		blah := phone[int(v-'1')]
		for _, vi := range blah {
			res[k] += string(vi)

		}
	}
	return res
}
