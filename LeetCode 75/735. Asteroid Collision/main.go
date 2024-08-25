package main

import (
	"fmt"
	"math"
)

func main() {
	res := asteroidCollision([]int{10, 2, -5})
	fmt.Println(res)
}

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
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

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

func asteroidCollision(asteroids []int) []int {
	var stack Stack

	for _, asteroid := range asteroids {
		// Case 1: asteroid is positive number.
		if asteroid > 0 {
			stack.Push(asteroid)
		} else { // Case 2: asteroid is negative number.
			// Check => if a stack doesn't empty and the top element of stack are more than 0 and less than asteroid.
			for !stack.IsEmpty() && stack.Peek() > 0 && stack.Peek() < int(math.Abs(float64(asteroid))) {
				// Pop or Remove the top element of stack.
				stack.Pop()
			}

			// Check => if a stack does empty and the top element of stack less than 0.
			if stack.IsEmpty() || stack.Peek() < 0 {
				// Push the asteroid to stack.
				stack.Push(asteroid)
			} else if asteroid+stack.Peek() == 0 { // the sum of asteroid and the top element of stack is equal 0.
				// Pop or Remove the top element of stack.
				stack.Pop()
			}
		}
	}

	return stack.items
}
