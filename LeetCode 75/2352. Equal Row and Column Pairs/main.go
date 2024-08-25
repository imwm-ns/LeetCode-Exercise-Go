package main

import (
	"fmt"
	"strconv"
)

func main() {
	grid := [][]int{{11, 1}, {1, 11}}
	resp := equalPairs(grid)

	fmt.Println(resp)
}

func equalPairs(grid [][]int) int {
	rows := make([]string, len(grid))
	cols := make([]string, len(grid))
	_ = cols

	existMap := make(map[string]int)

	for i, itm := range grid {
		newItem := ""

		for _, el := range itm {
			newItem += strconv.Itoa(el) + ","
		}

		rows[i] = newItem
		existMap[newItem]++
	}

	fmt.Println(rows)
	fmt.Println(existMap)

	for i := 0; i < len(grid); i++ {
		newItem := ""
		for j := 0; j < len(grid); j++ {
			newItem += strconv.Itoa(grid[j][i]) + ","
		}
		cols[i] = newItem
	}

	resp := 0

	for _, col := range cols {
		if v, ok := existMap[col]; ok {
			resp += v
		}
	}

	return resp
}
