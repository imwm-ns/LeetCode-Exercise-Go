package main

import "fmt"

func main() {
	var root *TreeNode = &TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 9,
	}
	root.Right = &TreeNode{
		Val: 20,
	}
	root.Right.Left = &TreeNode{
		Val: 15,
	}
	root.Right.Right = &TreeNode{
		Val: 7,
	}

	res := maxDepth(root)

	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}
