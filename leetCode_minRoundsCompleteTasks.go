package main

import "fmt"

func main() {
	t1 := []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}
	t2 := []int{2, 3, 3}
	t3 := []int{2, 2, 2, 3, 3, 5, 5, 5, 5, 5, 5, 5}

	for i := 1; i < 200; i++ {
		if i%2 != 0 && i%3 != 0 && (i-2)%3 != 0 && (i-3)%2 != 0 {
			fmt.Println("what kinda number are you, ", i)
		}
	}

	fmt.Println(t1, " takes ", minimumRounds(t1), " rounds of work to finish")
	fmt.Println(t2, " takes ", minimumRounds(t2), " rounds of work to finish")
	fmt.Println(t3, " takes ", minimumRounds(t3), " rounds of work to finish")
}

func minimumRounds(tasks []int) int {
	var taskData map[int]int = make(map[int]int) // key = difficulty, value= freq
	total := 0
	for _, d := range tasks {
		//		_, ok := taskData[d]
		//		if ok {
		taskData[d]++
		//		} else {
		//		taskData[d] = 1
		//		}
	}
	for _, freq := range taskData {
		if freq == 1 {
			return -1
		}
		if freq%3 == 0 {
			total += freq / 3
		} else if freq%3 == 2 {
			total += 1 + (freq / 3)
		} else if freq%3 == 1 {
			total += 2 + ((freq - 2) / 3)
		}
	}
	return total
}
