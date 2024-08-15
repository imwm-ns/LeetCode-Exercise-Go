package main

func main() {
	arr := []int{1, 2}
	uniqueOccurrences(arr)
}

func uniqueOccurrences(arr []int) bool {
	mp := make(map[int]int)
	seen := make(map[int]bool)

	for _, v := range arr {
		mp[v]++
	}

	for _, v := range mp {
		if _, ok := seen[v]; ok {
			return false
		}
		seen[v] = true
	}

	return true
}
