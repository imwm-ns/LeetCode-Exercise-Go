package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}

	fmt.Println(pivotIndex(nums))
}

func pivotIndex(nums []int) int {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0

	for i, v := range nums {
		if leftSum == (totalSum - leftSum - v) {
			return i
		}
		leftSum += v
	}

	return -1
}
