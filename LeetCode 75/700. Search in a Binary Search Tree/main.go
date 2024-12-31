package main

func main() {
	var root *TreeNode = &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	res := searchBST(root, 2)

	println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val == val:
		return root
	case root.Val < val:
		return searchBST(root.Right, val)
	case root.Val > val:
		return searchBST(root.Left, val)
	}
	return nil
}
