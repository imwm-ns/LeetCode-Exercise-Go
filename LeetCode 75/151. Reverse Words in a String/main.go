package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("a good   example")
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	formattedStr := strings.Join(strings.Fields(s), " ")
	words := strings.Split(formattedStr, " ")

	left, right := 0, len(words)-1

	for left < right {
		words[left], words[right] = words[right], words[left]
		left++
		right--
	}

	return strings.Join(words, " ")
}
