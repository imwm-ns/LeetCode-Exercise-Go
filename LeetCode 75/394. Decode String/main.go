package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))
}

type Stack struct {
	items []string
}

func (s *Stack) Push(data string) {
	s.items = append(s.items, data)
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack) Peek() string {
	return s.items[len(s.items)-1]
}

func (s *Stack) Length() int {
	return len(s.items)
}

func decodeString(s string) string {
	var stack Stack
	var sb strings.Builder
	for _, char := range s {
		if char != ']' {
			stack.Push(string(char))
		} else {
			var subStr string

			for stack.Peek() != "[" {
				subStr = stack.Peek() + subStr
				stack.Pop()
			}

			stack.Pop()

			cnt := ""

			for stack.Length() > 0 && stack.Peek() >= "0" && stack.Peek() <= "9" {
				cnt = stack.Peek() + cnt
				stack.Pop()
			}

			num, _ := strconv.Atoi(cnt)

			repeated := strings.Repeat(subStr, num)
			stack.Push(repeated)
		}
	}

	for _, itm := range stack.items {
		sb.WriteString(itm)
	}

	return sb.String()
}
