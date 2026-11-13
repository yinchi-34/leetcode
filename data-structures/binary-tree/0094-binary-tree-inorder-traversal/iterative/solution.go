package solution

// TreeNode is the binary tree node defined by LeetCode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderTraversal returns the values of the tree rooted at root in inorder
// (left subtree, root, right subtree), using an explicit stack of the
// ancestors whose own values have not been visited yet.
func InorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	node := root

	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, tmp.Val)
			// 先访问完当前节点的左右节点
			node = tmp.Right
		}
	}
	return res
}
