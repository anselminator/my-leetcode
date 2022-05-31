package main

import "fmt"

// https://leetcode.com/problems/string-to-integer-atoi/
func main() {
	var (
		a = "1"
		b = "-1"
		c = "   42"
		d = "   -42   3 323e32"
		e = "  292131209    "
		f = "   4   44343535  4353  5345 345 "
		g = "+32"
		h = "+9000"
		i = "-90001"
		j = "9223372036854775808"
		k = "-9223372036854775809"
		l = "18446744073709551617"
		m = "10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"
		n = "-000000000000000000000000000000000000000000000000001"
	)
	fmt.Println(pow(2, 4))
	fmt.Println(pow(4, 2))
	fmt.Println(pow(10, 3))
	fmt.Println(pow(10, 0))

	fmt.Println("string", a, "is int", myAtoi(a))
	fmt.Println("string", b, "is int", myAtoi(b))
	fmt.Println("string", c, "is int", myAtoi(c))
	fmt.Println("string", d, "is int", myAtoi(d))
	fmt.Println("string", e, "is int", myAtoi(e))
	fmt.Println("string", f, "is int", myAtoi(f))
	fmt.Println("string", g, "is int", myAtoi(g))
	fmt.Println("string", h, "is int", myAtoi(h))
	fmt.Println("string", i, "is int", myAtoi(i))
	fmt.Println("string", j, "is int", myAtoi(j))
	fmt.Println("string", k, "is int", myAtoi(k))
	fmt.Println("string", l, "is int", myAtoi(l))
	fmt.Println("string", m, "is int", myAtoi(m))
	fmt.Println("string", n, "is int", myAtoi(n))

}

func myAtoi(s string) int {
	var res int64
	var res2 int
	var sign bool = true
	var il int
	var r2 rune
	var leave bool = false
	if s == "" {
		return 0
	}
	//remove leading whitespace
	for _, r := range s {
		switch r {
		case ' ':
			s = s[1:]
			continue
		case '+':
			sign = true
			s = s[1:]
			leave = true
			break
		case '-':
			sign = false
			s = s[1:]
			leave = true
			break
		default:
			leave = true
			break
		}
		if leave {
			break
		}
	}
	leave = false

	//remove leadinng zeros
	for _, r := range s {
		if r == '0' {
			s = s[1:]
		} else {
			break
		}

	}

	for il, r2 = range s {
		switch r2 {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			il--
			leave = true
			break
		}
		if leave {
			break
		}
	}
	if s == "" {
		return 0
	}

	//remove any trailing non digit characters
	s = s[:il+1]
	fmt.Println("fuck(", s, ")")

	for j := 1; j <= len(s); j++ {
		res += int64(s[len(s)-j]-'0') * int64(pow(10, j-1))
		if res > 2147483647 || j > 32 {
			if sign {
				return 2147483647
			} else {
				return -2147483648
			}
		}

	}
	fmt.Println(sign, res)
	if res < 0 { // overflow
		if sign {
			return 2147483647
		} else {
			return -2147483648
		}
	}
	if res > 0 && !sign {
		res = -res
	}
	if res > 2147483647 {
		res = 2147483647
	}
	if res < -2147483648 {
		res = -2147483648
	}

	res2 = int(res)
	return res2
}

func pow(n int, m int) int {
	if m == 0 {
		return 1
	}
	r := n
	for i := 1; i < m; i++ {
		r = r * n
	}
	return r
}
