package solution

// TreeNode 是 LeetCode 给出的二叉树结点定义。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 右侧两个递归调用都求值完毕后才同时赋值，所以不需要临时变量。
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
