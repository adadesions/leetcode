package twoSum

func TwoSum(nums []int, target int) [2]int {
	ans := [2]int{}

	for i, n := range nums {
		for j, m := range nums {
			if i == j {
				continue
			}

			if n+m == target {
				if i > j {
					ans = [2]int{j, i}
				} else {
					ans = [2]int{i, j}
				}
			}
		}
	}
	return ans
}
