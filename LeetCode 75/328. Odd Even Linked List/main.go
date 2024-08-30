package main

import "fmt"

func main() {
	var head *ListNode = &ListNode{
		Val: 1,
	}
	head.Next = &ListNode{
		Val: 2,
	}
	head.Next.Next = &ListNode{
		Val: 3,
	}
	head.Next.Next.Next = &ListNode{
		Val: 4,
	}
	head.Next.Next.Next.Next = &ListNode{
		Val: 5,
	}

	head = oddEvenList(head)

	printLinkList(head)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var odd, even *ListNode = head, head.Next
	var evenHead *ListNode = even

	for (odd.Next != nil) && (even.Next != nil) {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next

		odd = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}

func printLinkList(head *ListNode) {
	var temp *ListNode = head

	for temp != nil {
		fmt.Printf("%d ", temp.Val)
		temp = temp.Next
	}
	fmt.Println("")
}
