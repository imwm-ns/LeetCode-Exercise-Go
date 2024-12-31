package main

func main() {
	var root *TreeNode = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}

	result := rightSideView(root)

	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	rightSide := make([]int, 0)

	bfs(root, 0, &rightSide)

	return rightSide
}

func bfs(node *TreeNode, level int, queue *[]int) {
	if node == nil {
		return
	}

	if len(*queue) <= level {
		*queue = append(*queue, node.Val)
	} else {
		(*queue)[level] = node.Val
	}

	bfs(node.Left, level+1, queue)
	bfs(node.Right, level+1, queue)
}
