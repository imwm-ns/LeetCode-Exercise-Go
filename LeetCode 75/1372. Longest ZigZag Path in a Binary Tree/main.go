package main

func main() {
	var root *TreeNode = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Left = &TreeNode{Val: 1}
	root.Right.Right.Right = &TreeNode{Val: 1}

	result := longestZigZag(root)

	println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	return max(dfs(root.Left, true, 0), dfs(root.Right, false, 0))
}

func dfs(node *TreeNode, isLeft bool, counter int) int {
	if node == nil {
		return counter
	}

	if isLeft {
		return max(dfs(node.Left, true, 0), dfs(node.Right, false, counter+1))
	}

	return max(dfs(node.Left, true, counter+1), dfs(node.Right, false, 0))
}
