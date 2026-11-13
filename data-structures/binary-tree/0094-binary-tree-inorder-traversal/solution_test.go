package solution

import (
	"slices"
	"testing"

	iterative "leetcode-go/data-structures/binary-tree/0094-binary-tree-inorder-traversal/iterative"
	recursive "leetcode-go/data-structures/binary-tree/0094-binary-tree-inorder-traversal/recursive"
)

// TreeNode is a test-local tree. Each solution package defines its own TreeNode,
// as LeetCode does, so tests build this tree once and convert it per package.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func toRecursive(t *TreeNode) *recursive.TreeNode {
	if t == nil {
		return nil
	}
	return &recursive.TreeNode{Val: t.Val, Left: toRecursive(t.Left), Right: toRecursive(t.Right)}
}

func toIterative(t *TreeNode) *iterative.TreeNode {
	if t == nil {
		return nil
	}
	return &iterative.TreeNode{Val: t.Val, Left: toIterative(t.Left), Right: toIterative(t.Right)}
}

var solutions = []struct {
	name string
	fn   func(*TreeNode) []int
}{
	{"recursive", func(t *TreeNode) []int { return recursive.InorderTraversal(toRecursive(t)) }},
	{"iterative", func(t *TreeNode) []int { return iterative.InorderTraversal(toIterative(t)) }},
}

// buildTree builds a tree from LeetCode's level-order form, where nil marks a missing child.
func buildTree(vals []any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     []any
		expected []int
	}{
		{
			name:     "Example 1",
			root:     []any{1, nil, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "Example 2",
			root:     []any{1, 2, 3, 4, 5, nil, 8, nil, nil, 6, 7, 9},
			expected: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name:     "Empty tree",
			root:     []any{},
			expected: []int{},
		},
		{
			name:     "Single node",
			root:     []any{1},
			expected: []int{1},
		},
		{
			name:     "Left-skewed chain",
			root:     []any{3, 2, nil, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Binary search tree gives sorted order",
			root:     []any{4, 2, 6, 1, 3, 5, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Negative values",
			root:     []any{0, -1, 1},
			expected: []int{-1, 0, 1},
		},
	}

	for _, sol := range solutions {
		for _, test := range tests {
			t.Run(sol.name+"/"+test.name, func(t *testing.T) {
				result := sol.fn(buildTree(test.root))
				if !slices.Equal(result, test.expected) {
					t.Errorf("%s(%v) = %v, want %v", sol.name, test.root, result, test.expected)
				}
			})
		}
	}
}

// TestInorderTraversalDeepChain checks both solutions on a 10^4-node right-skewed chain,
// where the explicit stack never holds more than one node.
func TestInorderTraversalDeepChain(t *testing.T) {
	const n = 10000
	var root *TreeNode
	for i := n; i >= 1; i-- {
		root = &TreeNode{Val: i, Right: root}
	}
	expected := make([]int, n)
	for i := range expected {
		expected[i] = i + 1
	}
	for _, sol := range solutions {
		t.Run(sol.name, func(t *testing.T) {
			if result := sol.fn(root); !slices.Equal(result, expected) {
				t.Errorf("%s on a %d-node chain returned a wrong order", sol.name, n)
			}
		})
	}
}
