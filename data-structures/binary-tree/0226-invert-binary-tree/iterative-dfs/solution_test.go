package solution

import (
	"slices"
	"testing"
	"time"
)

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

// levelOrder 是 buildTree 的逆过程：按 LeetCode 的层序格式输出，并去掉末尾的 null。
// 若同一个结点被访问两次（例如左右孩子指向同一棵子树），结果已不是一棵树，直接判失败，
// 以免在这种结构上展开出指数长的输出。
func levelOrder(t *testing.T, root *TreeNode) []int {
	t.Helper()
	res := []int{}
	seen := map[*TreeNode]bool{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			res = append(res, null)
			continue
		}
		if seen[node] {
			t.Fatalf("node %d is reachable twice; the result is not a tree", node.Val)
		}
		seen[node] = true
		res = append(res, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(res) > 0 && res[len(res)-1] == null {
		res = res[:len(res)-1]
	}
	return res
}

// chain 返回 n 个结点连成的单链：left 为 true 时全部挂在左孩子上，否则全部挂在右孩子上。
// 结点值从根往下依次为 0, 1, ..., n-1。
func chain(n int, left bool) *TreeNode {
	var root *TreeNode
	for i := n - 1; i >= 0; i-- {
		node := &TreeNode{Val: i}
		if left {
			node.Left = root
		} else {
			node.Right = root
		}
		root = node
	}
	return root
}

// isChain 检查 root 是否恰好是 n 个结点、值依次为 0..n-1、全部挂在指定一侧的单链。
func isChain(root *TreeNode, n int, left bool) bool {
	node := root
	for i := 0; i < n; i++ {
		if node == nil || node.Val != i {
			return false
		}
		if left {
			if node.Right != nil {
				return false
			}
			node = node.Left
		} else {
			if node.Left != nil {
				return false
			}
			node = node.Right
		}
	}
	return node == nil
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name     string
		root     []int
		expected []int
	}{
		{
			name:     "example_1",
			root:     []int{4, 2, 7, 1, 3, 6, 9},
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "example_2",
			root:     []int{2, 1, 3},
			expected: []int{2, 3, 1},
		},
		{
			name:     "example_3_empty",
			root:     []int{},
			expected: []int{},
		},
		{
			name:     "single_node",
			root:     []int{1},
			expected: []int{1},
		},
		{
			// 只有左孩子：翻转后应变成只有右孩子
			name:     "left_child_only",
			root:     []int{1, 2},
			expected: []int{1, null, 2},
		},
		{
			name:     "right_child_only",
			root:     []int{1, null, 2},
			expected: []int{1, 2},
		},
		{
			// 只交换根的两个孩子而不递归时，第二层以下保持原样，会得到 [1,3,2,6,7,4,5]
			name:     "deeper_levels_must_swap",
			root:     []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 3, 2, 7, 6, 5, 4},
		},
		{
			// 不对称、带空位的树
			name:     "asymmetric_with_gaps",
			root:     []int{1, 2, 3, null, 4, 5, null, null, null, null, 6},
			expected: []int{1, 3, 2, null, 5, 4, null, 6},
		},
		{
			name:     "negative_and_duplicate_values",
			root:     []int{0, -100, 100, -100, null, null, 100},
			expected: []int{0, 100, -100, 100, null, null, -100},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			root := buildTree(test.root)
			result := invertTree(root)
			if result != root {
				t.Errorf("invertTree(%v) returned a different root pointer", test.root)
			}
			if got := levelOrder(t, result); !slices.Equal(got, test.expected) {
				t.Errorf("invertTree(%v) = %v, want %v", test.root, got, test.expected)
			}
		})
	}
}

// TestInvertTreeChain 检查结点数取上限 100 且退化成单链时，左链翻成右链、右链翻成左链。
// 加一个超时，让“同一个结点被反复处理”之类的错误表现为失败而不是卡住。
func TestInvertTreeChain(t *testing.T) {
	const n = 100
	for _, left := range []bool{true, false} {
		done := make(chan *TreeNode, 1)
		go func() { done <- invertTree(chain(n, left)) }()
		select {
		case root := <-done:
			if !isChain(root, n, !left) {
				t.Errorf("invertTree(chain(%d, left=%v)) is not the mirrored chain", n, left)
			}
		case <-time.After(time.Second):
			t.Fatalf("invertTree(chain(%d, left=%v)) did not finish within 1s", n, left)
		}
	}
}

// TestInvertTreeTwice 检查翻转是对合：连续翻转两次得到原树。
func TestInvertTreeTwice(t *testing.T) {
	inputs := [][]int{
		{},
		{1},
		{4, 2, 7, 1, 3, 6, 9},
		{1, 2, 3, null, 4, 5, null, null, null, null, 6},
	}
	for _, vals := range inputs {
		if got := levelOrder(t, invertTree(invertTree(buildTree(vals)))); !slices.Equal(got, vals) {
			t.Errorf("invertTree twice on %v = %v, want the original", vals, got)
		}
	}
}
