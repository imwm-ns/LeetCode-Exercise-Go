package main

import "fmt"

func main() {
	nums1, nums2 := []int{1, 2, 3}, []int{2, 4, 6}
	fmt.Println(findDifference(nums1, nums2))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	a, b := unique(nums1), unique(nums2)

	ans := make([][]int, 0)

	val1 := isNotInTarget(a, b)
	val2 := isNotInTarget(b, a)

	ans = append(ans, val1, val2)

	return ans
}

func unique(nums []int) map[int]bool {
	mp := make(map[int]bool)

	for _, v := range nums {
		mp[v] = true
	}

	return mp
}

func isNotInTarget(a, b map[int]bool) []int {
	ans := make([]int, 0)

	for k, v := range a {
		_ = v
		if _, ok := b[k]; !ok {
			ans = append(ans, k)
		}
	}

	return ans
}
