package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	switch {
	case root == nil:
		return nil
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	case root.Left == nil:
		return root.Right
	case root.Right == nil:
		return root.Left
	default:
		node := root.Right

		for node.Left != nil {
			node = node.Left
		}

		root.Val = node.Val
		root.Right = deleteNode(root.Right, node.Val)
	}
	return root
}
