package main

import "fmt"

func main() {
	r1 := "UUURRRDDDDDDLLLLLLUUURRR"
	r2 := "UDLDLLDLDURURURURURURURUURURU"
	r3 := "UDLR"
	fmt.Println("move ROBOT")
	fmt.Println(r1, judgeCircle(r1))
	fmt.Println(r2, judgeCircle(r2))
	fmt.Println(r3, judgeCircle(r3))
	//	fmt.Println(judgeCircle(r4))
}

func judgeCircle(moves string) bool {
	var x int = 0
	var y int = 0
	for _, r := range moves {
		switch string(r) {
		case "U":
			y++
		case "D":
			y--
		case "L":
			x--
		case "R":
			x++
		}
	}
	return x == 0 && y == 0
}
