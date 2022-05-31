package main

import "fmt"

func main() {
	board1 := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	fmt.Println("Conway\\'s Game of Life, in Go, from leetcode")
	fmt.Println(board1)
}

func gameOfLife(board [][]int) {
	var sum int = 0
	var m int = len(board)
	var n int = len(board[0])
	var nextState [][]int

	nextState[sum][0] = board[1][1]
	return
}
