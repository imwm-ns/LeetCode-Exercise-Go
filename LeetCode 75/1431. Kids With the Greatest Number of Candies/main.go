package main

import (
	"fmt"
	"slices"
)

func main() {
	result := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	fmt.Println(result)
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := slices.Max(candies)

	result := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		var res bool
		if candies[i]+extraCandies >= max {
			res = true
		} else {
			res = false
		}

		result[i] = res
	}

	return result
}
