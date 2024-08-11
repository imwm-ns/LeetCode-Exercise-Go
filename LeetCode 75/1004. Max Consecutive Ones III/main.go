package main

import "fmt"

func main() {
	nums, k := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2
	fmt.Println(longestOnes(nums, k))
}

func longestOnes(nums []int, k int) int {
	i, j, zero, mx := 0, 0, 0, 0

	for j < len(nums) {
		if nums[j] == 0 {
			zero++
		}
		for zero > k {
			if nums[i] == 0 {
				zero--
			}
			i++
		}
		j++
		mx = max(mx, j-i)
	}

	return mx
}
