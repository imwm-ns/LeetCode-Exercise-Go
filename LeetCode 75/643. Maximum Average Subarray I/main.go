package main

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{2, 1, 5, 1, 3, 2}, 3))
}

func findMaxAverage(nums []int, k int) float64 {
	// ** Using Sliding Window Algorithm for solve the problem. ** //

	// Base Condition
	if k > len(nums) {
		return .0
	}

	// Initialize the variable that is slidingWindow to store the sum value of elements in slice.
	slidingWindow := 0

	// Compute the sum of the first 'k' elements.
	for i := 0; i < k; i++ {
		slidingWindow += nums[i]
	}

	// Determine maxSum variable with slidingWindow to store the sum of the first 'k elements.
	maxSum := slidingWindow

	// Iterate the element for slide the window over the array and updating the value of maxSum
	for i := k; i < len(nums); i++ {
		// Slide the window over the array by increase the value of next index and subtract the leftmost index of window.
		slidingWindow += nums[i] - nums[i-k]

		// Create the condition for check if the value of slidingWindow more than the value of maxSum,
		// Then updating the value of maxSum with slidingWindow.
		if slidingWindow > maxSum {
			maxSum = slidingWindow
		}
	}

	// Return the average of maximum sum of the target array.
	return float64(maxSum) / float64(k)
}
