package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(s string) string {
	start, end := 0, len(s)-1
	vowels := "aeiouAEIOU"
	runeSlices := []rune(s)

	for start < end {
		for start < end &&
			!strings.Contains(vowels, string(runeSlices[start])) {
			start++
		}

		for start < end &&
			!strings.Contains(vowels, string(runeSlices[end])) {
			end--
		}

		if start < end {
			runeSlices[start], runeSlices[end] = runeSlices[end], runeSlices[start]
			start++
			end--
		}
	}

	return string(runeSlices)
}
