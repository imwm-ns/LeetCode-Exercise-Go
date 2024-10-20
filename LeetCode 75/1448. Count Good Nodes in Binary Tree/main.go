package main

import "fmt"

func main() {
	var root *TreeNode = &TreeNode{
		Val: 3,
	}
	root.Left = &TreeNode{
		Val: 3,
	}
	root.Left.Left = &TreeNode{
		Val: 4,
	}
	root.Left.Right = &TreeNode{
		Val: 2,
	}

	fmt.Println(goodNodes(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isGood(node *TreeNode, mx int) int {
	goodCounter := 0

	if node == nil {
		return 0
	}

	if node.Val >= mx {
		goodCounter = 1
		mx = node.Val
	}

	return goodCounter + isGood(node.Left, mx) + isGood(node.Right, mx)
}

func goodNodes(root *TreeNode) int {
	return isGood(root, root.Val)
}
