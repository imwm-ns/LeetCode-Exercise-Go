package main

func main() {

}

func longestSubArray(nums []int) int {
	l, r, cnt, mx := 0, 0, 0, 0

	for r < len(nums) {
		if nums[r] == 0 {
			cnt++
		}

		for cnt > 1 {
			if nums[l] == 0 {
				cnt--
			}
			l++
		}

		mx = max(mx, r-l)
		r++
	}
	return mx
}
