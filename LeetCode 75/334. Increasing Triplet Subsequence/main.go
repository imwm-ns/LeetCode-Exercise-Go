package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	a, b := math.MaxInt64, math.MaxInt64

	for _, val := range nums {
		if val <= a {
			a = val
		} else if val <= b {
			b = val
		} else {
			return true
		}
	}
	return false
}
