package solution

import "testing"

// null 表示层序数组中的空位置。题目保证结点值在 [-100, 100] 内，不会与它冲突。
const null = -1 << 31

// buildTree 按 LeetCode 的层序格式建树：依次为队列中每个结点读取左、右孩子，
// 遇到 null 表示该孩子为空；末尾省略的 null 不需要写出。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := queue[0]
		queue = queue[1:]
		if vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// chain 返回 n 个结点连成的单链：left 为 true 时全部挂在左孩子上，否则全部挂在右孩子上。
func chain(n int, left bool) *TreeNode {
	var root *TreeNode
	for i := 0; i < n; i++ {
		node := &TreeNode{Val: i%201 - 100}
		if left {
			node.Left = root
		} else {
			node.Right = root
		}
		root = node
	}
	return root
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected int
	}{
		{
			name:     "example_1",
			root:     buildTree([]int{3, 9, 20, null, null, 15, 7}),
			expected: 3,
		},
		{
			// 只有右孩子：若只看左子树会得到 1
			name:     "example_2_right_only",
			root:     buildTree([]int{1, null, 2}),
			expected: 2,
		},
		{
			name:     "empty_tree",
			root:     nil,
			expected: 0,
		},
		{
			name:     "single_node",
			root:     buildTree([]int{0}),
			expected: 1,
		},
		{
			// 左子树更深：若只看右子树会得到 2
			name:     "left_deeper",
			root:     buildTree([]int{1, 2, 3, 4, null, null, null, 5}),
			expected: 4,
		},
		{
			// 一侧子树为空、另一侧很深：若把 max 写成 min 会得到 1
			name:     "one_side_empty",
			root:     buildTree([]int{1, null, 2, null, 3, null, 4}),
			expected: 4,
		},
		{
			name:     "perfect_tree_depth_3",
			root:     buildTree([]int{1, 2, 3, 4, 5, 6, 7}),
			expected: 3,
		},
		{
			// 结点数取上限 10^4 且退化成单链：递归深度达到 10^4
			name:     "left_chain_max_nodes",
			root:     chain(10_000, true),
			expected: 10_000,
		},
		{
			name:     "right_chain_max_nodes",
			root:     chain(10_000, false),
			expected: 10_000,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := maxDepth(test.root)
			if result != test.expected {
				t.Errorf("maxDepth(%s) = %d, want %d", test.name, result, test.expected)
			}
		})
	}
}
