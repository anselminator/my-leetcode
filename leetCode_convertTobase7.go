package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 1
	b := 7
	c := 8
	d := 17890
	e := -8

	fmt.Println(a, " in base 7 = ", convertToBase7(a))
	fmt.Println(b, " in base 7 = ", convertToBase7(b))
	fmt.Println(c, " in base 7 = ", convertToBase7(c))
	fmt.Println(d, " in base 7 = ", convertToBase7(d))
	fmt.Println(e, " in base 7 = ", convertToBase7(e))
	fmt.Println(e, " in base 7 = ", convertToBase7(0))
	fmt.Println(a, " in base 4 = ", convertToBaseN(a, 4))
	fmt.Println(b, " in base 2 = ", convertToBaseN(b, 2))
	fmt.Println(c, " in base 8 = ", convertToBaseN(c, 8))
	fmt.Println(d, " in base 9 = ", convertToBaseN(d, 9))
	fmt.Println(e, " 16 in base 16 = ", convertToBaseN(16, 16))
	fmt.Println(e, " 64 in base 16 = ", convertToBaseN(64, 16))
	fmt.Println(e, " 100 in base 100 = ", convertToBaseN(100, 100))

}

func convertToBase7(num int) string {
	var r string = ""
	var sign string = ""
	if num == 0 {
		return "0"
	}
	if num < 0 {
		sign = "-"
		num = -num
	}
	for num != 0 {
		r = strconv.Itoa(num%7) + r
		num = num / 7
	}
	return sign + r
}
func convertToBaseN(num int, base int) string {
	var r string = ""
	var sign string = ""
	if num == 0 {
		return "0"
	}
	if num < 0 {
		sign = "-"
		num = -num
	}
	for num != 0 {
		r = strconv.Itoa(num%base) + r
		num = num / base
	}
	return sign + r
}
