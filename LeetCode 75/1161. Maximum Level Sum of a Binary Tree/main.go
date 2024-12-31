package main

import "math"

func main() {
	var root *TreeNode = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: -8}

	res := maxLevelSum(root)

	println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	hash := make(map[int]int)
	bfs(root, &hash, 0)

	level := 0
	maxVal := math.MinInt64

	for key, value := range hash {
		if value > maxVal {
			maxVal = value
			level = key
		}
	}

	return level + 1
}

func bfs(node *TreeNode, hash *map[int]int, level int) {
	if node == nil {
		return
	}

	(*hash)[level] += node.Val

	bfs(node.Left, hash, level+1)
	bfs(node.Right, hash, level+1)
}
