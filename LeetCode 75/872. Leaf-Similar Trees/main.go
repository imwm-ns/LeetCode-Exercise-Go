package main

import (
	"fmt"
	"slices"
)

func main() {
	// Unit Test

	// Defined the value of root1
	var root1 *TreeNode = &TreeNode{
		Val: 1,
	}
	root1.Left = &TreeNode{
		Val: 2,
	}
	root1.Right = &TreeNode{
		Val: 3,
	}

	// Defined the value of root2
	var root2 *TreeNode = &TreeNode{
		Val: 1,
	}
	root2.Left = &TreeNode{
		Val: 3,
	}
	root2.Right = &TreeNode{
		Val: 2,
	}

	fmt.Println(leafSimilar(root1, root2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func listDFS(node *TreeNode) []int {
	if node == nil {
		return []int{}
	}

	if node.Left == nil && node.Right == nil {
		return []int{node.Val}
	}

	return append(listDFS(node.Left), listDFS(node.Right)...)
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	return slices.Equal(listDFS(root1), listDFS(root2))
}
