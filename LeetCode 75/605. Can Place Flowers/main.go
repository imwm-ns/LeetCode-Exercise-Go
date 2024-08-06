package main

import "fmt"

func main() {
	result := canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)
	fmt.Println(result)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i, plot := range flowerbed {
		if plot == 1 {
			continue
		}
		if (i == 0 || flowerbed[i-1] == 0) &&
			(len(flowerbed)-1 == i || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
	}

	return n <= 0
}
