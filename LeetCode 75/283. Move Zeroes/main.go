package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 3, 12}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(nums []int) {
	idx := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}

	for i := idx; i < len(nums); i++ {
		nums[i] = 0
	}
}

/*
	Debugging

	round 1: i = 0, v = 0 => found 0 then increase swpIdx

	round 2: i = 1, v = 1 => v is not equal 0 then go to (swpIdx > 0) condition => temp = 1, nums[1] = 0, n[1 - 1] = temp(1)
	resp => [1, 0, 0, 3, 12]

	round 3: i = 2, v = 0 => found 0 then increase swpIdx

	round 4: i = 3, v = 3 => v is not equal 0 then go to (swpIdx > 0) condition => temp = 3, num[3] = 0, n[3 - 2] = temp(3)
	resp => [1, 3, 0, 0, 12]

	round 5: i = 4, v = 12 => v is not equal 0 then go to (swpIdx > 0) condition => temp = 12, num[4] = 0, n[4 - 2] = temp(12)
	resp => [1, 3, 12, 0, 0]

*/

func moveZeroes1(nums []int) {
	swpIdx := 0

	for i, v := range nums {
		if v == 0 {
			swpIdx++
		} else if swpIdx > 0 {
			temp := v
			nums[i] = 0
			nums[i-swpIdx] = temp
		}
	}
}
