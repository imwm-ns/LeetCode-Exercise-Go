package main

import "fmt"

func main() {
	fmt.Println(maxOperations([]int{3, 1, 3, 4, 3}, 6))
}

func maxOperations(nums []int, k int) int {
	ans := 0

	hmp := make(map[int]int)

	if len(nums) <= 1 {
		return ans
	}

	for _, num := range nums {
		if num > k {
			continue
		}

		diff := k - num

		if val, exist := hmp[diff]; exist && val > 0 {
			ans++
			hmp[diff]--
		} else {
			hmp[num]++
		}
	}

	return ans
}
