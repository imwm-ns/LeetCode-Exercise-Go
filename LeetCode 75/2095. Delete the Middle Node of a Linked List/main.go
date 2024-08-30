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

	head = deleteMiddle(head)

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

func deleteMiddle(head *ListNode) *ListNode {
	if (head == nil) || (head.Next == nil) {
		return nil
	}

	slow, fast := head, head.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

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
