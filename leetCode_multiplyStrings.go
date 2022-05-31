package main

import "fmt"

func main() {
	fmt.Println("multiplying two numbers as Strings, go, leetCode")
	fmt.Println("123123123+456", 123_123_123+456, " .... and with strings:", addStrings("123123123", "456"))
	fmt.Println("123+6=", 123+6, " .... and with strings:", addStrings("123", "6"))
	fmt.Println("123+456=", 123+456, " .... and with strings:", addStrings("123", "456"))
	fmt.Println("456+77=", 456+77, " .... and with strings:", addStrings("456", "77"))
	fmt.Println("1+111111110=", 1+111_111_110, " .... and with strings:", addStrings("1", "111111110"))
	fmt.Println("1+9=", 1+9, " .... and with strings:", addStrings("1", "9"))
	fmt.Println("1+999=", 1+999, " .... and with strings:", addStrings("1", "999"))
	fmt.Println("1+99999=", 1+99_999, " .... and with strings:", addStrings("1", "99999"))
}

func addStrings(num1 string, num2 string) string {
	var lnum1 int = len(num1)
	var lnum2 int = len(num2)
	var sum string = ""
	var smaller, bigger string
	var carry int = 0
	var dsum int = 0

	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}

	if lnum1 >= lnum2 {
		bigger = num1
		smaller = num2
		for i := 0; i < lnum1-lnum2; i++ {
			smaller = "0" + smaller
		}
	} else {
		bigger = num2
		smaller = num1
		for i := 0; i < lnum2-lnum1; i++ {
			smaller = "0" + smaller
		}
	}

	for j := len(bigger) - 1; j >= 0; j-- {
		dsum = (carry + int(bigger[j]-'0') + int(smaller[j]-'0')) % 10
		carry = (carry + int(bigger[j]-'0') + int(smaller[j]-'0')) / 10
		sum = string(dsum+'0') + sum
	}
	if carry == 1 {
		sum = "1" + sum
	}
	return sum
}

func multiply(num1 string, num2 string) string {

}
