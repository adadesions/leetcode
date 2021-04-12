package twoSum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   [2]int
	}{
		{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		{[]int{3, 2, 4}, 6, [2]int{1, 2}},
		{[]int{3, 3}, 6, [2]int{0, 1}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.nums)
		t.Run(testname, func(t *testing.T) {
			ans := TwoSum(tt.nums, tt.target)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}
