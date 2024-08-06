package main

import (
	"fmt"
	"strconv"
)

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	resp := compress(chars)
	fmt.Println(resp)
}

func compress(chars []byte) int {
	read, write := 0, 0

	for read < len(chars) {
		currentChar := chars[read]
		count := 0

		// Counting the number of consecutive repeating characters
		for read < len(chars) && chars[read] == currentChar {
			read++
			count++
		}

		// Writing the compression result to array
		chars[write] = currentChar
		write++

		// Check if the number of consecutive repeating characters more than 1
		if count > 1 {
			cnt := strconv.Itoa(count)
			for _, idx := range cnt {
				chars[write] = byte(idx)
				write++
			}
		}
	}

	return write
}
