package main

import "fmt"

type Stack struct {
	s    []int
	top  int
	size int
}

func main() {
	s1 := "()[]}"
	s2 := "()))"
	s3 := "{]"
	s4 := "{{}[][[([]]]}"
	s5 := "{{({})}}}"
	fmt.Println(s1, " isvalid = ", isValid(s1))
	fmt.Println(s2, " isvalid = ", isValid(s2))
	fmt.Println(s3, " isvalid = ", isValid(s3))
	fmt.Println(s4, " isvalid = ", isValid(s4))
	fmt.Println(s5, " isvalid = ", isValid(s5))
}

func isValid(s string) bool {
	stac := makeStack()
	for _, r := range s {
		switch r {
		case '(':
			push(1, stac)
		case '[':
			push(2, stac)
		case '{':
			push(3, stac)
		case ')':
			br, ok := pop(stac)
			if !ok || (br != 1) {
				return false
			}
		case ']':
			br, ok := pop(stac)
			if !ok || (br != 2) {
				return false
			}
		case '}':
			br, ok := pop(stac)
			if !ok || (br != 3) {
				return false
			}
		}
	}
	return stac.size == 0
}

func makeStack() *Stack {
	var s Stack = Stack{[]int{}, 0, 0}
	return &s
}
func push(v int, s *Stack) {
	s.s = append([]int{v}, s.s...)
	s.size++
}

func pop(s *Stack) (int, bool) {
	if s.size == 0 {
		return 0, false
	}
	v := s.s[0]
	s.s = s.s[1:]
	s.size--
	return v, true
}
