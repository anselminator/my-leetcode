package main

import "fmt"

var reverseTable map[uint32]uint32 = make(map[uint32]uint32)

// https://leetcode.com/problems/reverse-bits/
func main() {
	var a uint32 = 127
	var b uint32 = 964176192
	var c uint32 = 4294967293
	var d uint32 = 43261596
	var e uint32 = 0b0101010101010101010101111111111
	fmt.Printf("%0b -> %0b\n", a, reverseBits(a))
	fmt.Printf("%0b -> %0b\n", a, reverseBits(a))
	fmt.Printf("%0b -> %0b\n", a, reverseBits(a))
	fmt.Printf("%0b -> %0b\n", a, reverseBits(a))
	fmt.Printf("%0b -> %0b\n", a, reverseBits(a))
	fmt.Printf("%0b -> %0b\n", a, reverseBits(a))
	fmt.Printf("%d -> %d\n", b, reverseBits(b))
	fmt.Printf("%d -> %d\n", c, reverseBits(c))
	fmt.Printf("%d -> %d\n", d, reverseBits(d))
	fmt.Printf("%d -> %d\n", d, reverseBits(d))
	fmt.Printf("%d -> %d\n", d, reverseBits(d))
	fmt.Printf("%d -> %d\n", d, reverseBits(d))
	fmt.Printf("%d -> %d\n", d, reverseBits(d))
	fmt.Printf("%0b -> %0b\n", e, reverseBits(e))
	fmt.Printf("%0b -> %0b\n", e, reverseBits(e))
	//	fmt.Println(reverseTable)
}

func reverseBits(num uint32) uint32 {
	v, ok := reverseTable[num]
	if ok {
		return v
	}
	var res uint32 = 0
	var i uint32
	for i = 0; i < 32; i++ {
		if (num & (1 << i)) > 0 {
			res |= 1 << (31 - i)
		}
	}
	reverseTable[num] = res
	return res
}
