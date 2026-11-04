package solution

// TreeNode is the binary tree node defined by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal returns the values of the tree rooted at root in inorder
// (left subtree, root, right subtree). Each call returns a new slice that the
// caller concatenates.
func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return []int{}
	}

	res = append(res, InorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, InorderTraversal(root.Right)...)

	return res
}
