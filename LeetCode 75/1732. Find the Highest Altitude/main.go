package main

import "fmt"

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Println(largestAltitude(gain))
}

func largestAltitude(gain []int) int {
	sum := 0
	hightest := sum

	for _, point := range gain {
		sum += point

		hightest = max(sum, hightest)
	}

	return hightest
}
