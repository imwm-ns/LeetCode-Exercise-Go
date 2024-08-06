package main

import (
	"fmt"
)

func main() {
	result := mergeAlternately("abc", "pqr")
	fmt.Println(result)
}

func mergeAlternately(word1 string, word2 string) string {
	index := 0
	result := []rune{}
	for index < len(word1) || index < len(word2) {
		if index < len(word1) {
			result = append(result, []rune(word1)[index])
		}
		if index < len(word2) {
			result = append(result, []rune(word2)[index])
		}
		index++
	}
	return string(result)
}
