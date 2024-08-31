package main

import "fmt"

func main() {
	head := &ListNode{
		Val: 4,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next.Next = &ListNode{
		Val: 3,
	}
	result := pairSum(head)

	fmt.Println(result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	TODO: Find the maximum of twin sum
	Step1:
		Initialize variables
			1. hash map to store the value of index that mean the index is key and the value of index is value.
			2. length to store the length of single linked list.
	Step2:
		Define a variable for store the result after found the maximum of twin sum.
	Step3:
		1. Iterate with length divide 2 and plus 1 that mean we should to iterate with the half of the length of single linked list.
		2. find the sum of twin pair.
		3. check if the sum more than the result. then if the sum more than the result, we should to change the result with sum.
	Finally:
		Return the result of the question.
*/

func pairSum(head *ListNode) int {
	hp := make(map[int]int)
	length := 0

	for head != nil {
		hp[length] = head.Val
		length++

		head = head.Next
	}

	result := 0

	for i := 0; i < length/2+1; i++ {
		sum := hp[i] + hp[length-i-1]

		if result < sum {
			result = sum
		}
	}

	return result
}
