package main

import "fmt"

func main() {
	s, k := "abciiidef", 3
	fmt.Println(maxVowels(s, k))
}

func maxVowels(s string, k int) int {
	if len(s) < k {
		return 0
	}

	mx, cnt := 0, 0

	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			cnt++
		}

		if i >= k && isVowel(s[i-k]) {
			cnt--
		}
		mx = max(mx, cnt)
	}

	return mx
}

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
