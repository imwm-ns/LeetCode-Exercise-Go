package main

func main() {
	removeStars("leet**cod*e")
}

func removeStars(s string) string {
	var stack Stack

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			stack.Pop()
		} else {
			stack.Push(s[i])
		}
	}

	return string(stack.items)
}

type Stack struct {
	items []byte
}

func (s *Stack) Push(data byte) {
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
