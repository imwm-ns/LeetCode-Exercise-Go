package main

import "fmt"

func main() {

}

func printLinkList(head *ListNode) {
	var temp *ListNode = head

	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Println("")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// Base case: Check the current value or next value is null. We will be return head.
	if head == nil || head.Next == nil {
		return head
	}

	// Declare 2 variables to store null and head linked list
	var prev *ListNode
	var cur *ListNode = head

	for cur != nil {
		// Store the next current value in next variable
		next := cur.Next

		// Reverse the linked list
		cur.Next = prev

		// Change the pointer of prev to current
		prev = cur

		// Change the pointer of current to next until found a null
		cur = next
	}

	// Finally. Return prev linked list to the result of this question.
	return prev
}
