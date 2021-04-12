package main

import "fmt"

func findMaxConsecutiveOnes(data []int) int {
	counter := 0
	max := 0
	isOne := false
	for _, d := range data {
		isOne = (d == 1)
		if isOne {
			counter += 1
		} else {
			counter = 0
		}

		if max < counter {
			max = counter
		}
		
	}

	return max
}

func main() {
	data := []int{1, 1, 1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(data))
	
}