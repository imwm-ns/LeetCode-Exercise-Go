package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	resp := make([]int, n)

	prefix := 1
	for i := 0; i < n; i++ {
		resp[i] = prefix
		prefix *= nums[i]
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		resp[i] *= suffix
		suffix *= nums[i]
	}

	return resp
}
