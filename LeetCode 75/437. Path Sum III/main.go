package main

func main() {
	targetSum := 8

	var root *TreeNode = &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: -3}
	root.Right.Right = &TreeNode{Val: 11}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Right = &TreeNode{Val: -2}
	root.Left.Right.Right = &TreeNode{Val: 11}

	result := pathSum(root, targetSum)

	println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode, target int, path []int) int {
	counter := 0

	if node == nil {
		return counter
	}

	path = append(path, node.Val)

	sum := 0

	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]

		if sum == target {
			counter++
		}
	}

	return counter + dfs(node.Left, target, path) + dfs(node.Right, target, path)
}

func pathSum(root *TreeNode, targetSum int) int {
	return dfs(root, targetSum, make([]int, 0))
}
