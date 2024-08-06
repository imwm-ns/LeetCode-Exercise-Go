package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	l, r, mx := 0, len(height)-1, 0

	for l < r {
		w, h := r-l, min(height[l], height[r])

		mx = max(w*h, mx)

		if height[l] < height[r] {
			l++
		} else if height[l] > height[r] {
			r--
		} else {
			l++
			r--
		}
	}
	return mx
}
