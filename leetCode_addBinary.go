package main

import "fmt"

func main() {
	// work in Progress
	fmt.Println("adding two Strings in binary, go, leetCode")
	fmt.Print("1111+110 = ")
	fmt.Printf("%b", 0b1111+0b110)
	fmt.Println(" .... and with strings:", addBinary("1111", "110"))
	fmt.Print("1+1 = ")
	fmt.Printf("%b", 0b1+0b1)

	fmt.Println(" .... and with strings:", addBinary("1", "1"))

	fmt.Print("100+100 = ")
	fmt.Printf("%b", 0b100+0b100)
	fmt.Println(" .... and with strings:", addBinary("100", "100"))
	/*	fmt.Println("456+77=", 456+77, " .... and with strings:", addStrings("456", "77"))
		fmt.Println("1+111111110=", 1+111_111_110, " .... and with strings:", addStrings("1", "111111110"))
		fmt.Println("1+9=", 1+9, " .... and with strings:", addStrings("1", "9"))
		fmt.Println("1+999=", 1+999, " .... and with strings:", addStrings("1", "999"))
		fmt.Println("1+99999=", 1+99_999, " .... and with strings:", addStrings("1", "99999"))
	*/
}

func addBinary(a string, b string) string {
	var lnum1 int = len(a)
	var lnum2 int = len(b)
	var sum string = ""
	var smaller, bigger string
	var carry int = 0
	var dsum int = 0

	if a == "0" {
		return b
	}
	if b == "0" {
		return a
	}

	if lnum1 >= lnum2 {
		bigger = a
		smaller = b
		for i := 0; i < lnum1-lnum2; i++ {
			smaller = "0" + smaller
		}
	} else {
		bigger = b
		smaller = a
		for i := 0; i < lnum2-lnum1; i++ {
			smaller = "0" + smaller
		}
	}

	for j := len(bigger) - 1; j >= 0; j-- {
		dsum = (carry + int(bigger[j]-'0') + int(smaller[j]-'0')) % 2
		carry = (carry + int(bigger[j]-'0') + int(smaller[j]-'0')) / 2
		sum = string(dsum+'0') + sum
	}
	if carry == 1 {
		sum = "1" + sum
	}
	return sum
}
