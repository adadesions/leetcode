package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(nums))
}

func findNumbers(nums []int) int {
	evens := 0

	for _, n := range nums {
		s := strconv.Itoa(n)
		if len(s) % 2 == 0 {
			evens += 1
		}
	}
	return evens
}
