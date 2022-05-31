package main

import "fmt"

func main() {
	fmt.Println(myPow(2.0, 10))
	fmt.Println(myPow(10.0, 3))
	fmt.Println(myPow(2.1, 3))
	var f float64 = myPow(2.0, -2)
	fmt.Println(myPow(2.0, -2))
	fmt.Println(f)

	fmt.Println(myPow(2.0, 999))
	fmt.Println(myPow(0.0, 10))

}
func myPow(x float64, n int) float64 {
	var sign bool = true
	if n < 0 {
		sign = false
		n = -n
	}
	res := 1.0
	if x == 0 {
		return 0.0
	}
	if n == 0 {
		return 1.0
	}
	for n > 0 {
		res *= x
		n--
	}
	if sign {
		return res
	} else {
		return 1.0 / res
	}
}
