package main

import (
	"fmt"
	"sort"
)

func main() {
	word1, word2 := "cabbba", "abbccc"
	res := closeStrings(word1, word2)

	fmt.Println(res)
}

func closeStrings(word1 string, word2 string) bool {
	// Check if the lengths of word1 and word2 aren't equal
	if len(word1) != len(word2) {
		return false
	}

	// Declare 2 slices to store the frequency of characters for both words
	fq1 := make([]int, 26)
	fq2 := make([]int, 26)

	// Populate the frequency of characters of both words to slices
	for i := 0; i < len(word1); i++ {
		fq1[word1[i]-'a']++
		fq2[word2[i]-'a']++
	}

	// Compare between both words have a same the sets of characters
	for i := 0; i < len(fq1); i++ {
		if (fq1[i] > 0 && fq2[i] == 0) || (fq1[i] == 0 && fq2[i] > 0) {
			return false
		}
	}

	// Sort the sequence of frequency in ascending order
	sort.Ints(fq1)
	sort.Ints(fq2)

	// Compare between both words have a same the frequencies
	for i := 0; i < len(fq1); i++ {
		if fq1[i] != fq2[i] {
			return false
		}
	}

	return true
}
