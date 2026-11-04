package solution

import "testing"

// buildIntersecting 按 LeetCode 的评测格式构造两条链表：
// listA 全部建出；listB 只建前 skipB 个结点，然后接到 A 的第 skipA 个结点（0-based）上。
// skipA == len(listA) 表示不相交，此时 listB 全部单独建出。
// 返回两个表头、预期的相交结点，以及两条链表的全部结点（用于检查结构未被修改）。
func buildIntersecting(listA, listB []int, skipA, skipB int) (headA, headB, want *ListNode, nodesA, nodesB []*ListNode) {
	nodesA = buildNodes(listA)
	if skipA == len(listA) {
		nodesB = buildNodes(listB)
	} else {
		nodesB = append(buildNodes(listB[:skipB]), nodesA[skipA:]...)
		if skipB > 0 {
			nodesB[skipB-1].Next = nodesA[skipA]
		}
		want = nodesA[skipA]
	}
	return headOf(nodesA), headOf(nodesB), want, nodesA, nodesB
}

// buildNodes 把值序列建成一条链表，按顺序返回各结点。
func buildNodes(vals []int) []*ListNode {
	nodes := make([]*ListNode, len(vals))
	for i := len(vals) - 1; i >= 0; i-- {
		nodes[i] = &ListNode{Val: vals[i]}
		if i+1 < len(vals) {
			nodes[i].Next = nodes[i+1]
		}
	}
	return nodes
}

func headOf(nodes []*ListNode) *ListNode {
	if len(nodes) == 0 {
		return nil
	}
	return nodes[0]
}

// seq 返回 1, 2, ..., n 的值序列，n 取上限时用于构造长链表。
func seq(n int) []int {
	vals := make([]int, n)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}

// sameStructure 检查从 head 出发依次经过的结点恰好是 nodes，且以 nil 结尾。
func sameStructure(head *ListNode, nodes []*ListNode) bool {
	cur := head
	for _, node := range nodes {
		if cur != node {
			return false
		}
		cur = cur.Next
	}
	return cur == nil
}

func TestGetIntersectionNode(t *testing.T) {
	tests := []struct {
		name         string
		listA, listB []int
		skipA, skipB int
	}{
		{
			// 值 1 在两条链表的私有部分各出现一次，但不是同一个结点：必须按结点比较，不能按值比较
			name:  "example_1",
			listA: []int{4, 1, 8, 4, 5},
			listB: []int{5, 6, 1, 8, 4, 5},
			skipA: 2, skipB: 3,
		},
		{
			name:  "example_2",
			listA: []int{1, 9, 1, 2, 4},
			listB: []int{3, 2, 4},
			skipA: 3, skipB: 1,
		},
		{
			// 不相交且长度不同：两个指针在第二圈末尾同时到达 nil
			name:  "example_3_no_intersection",
			listA: []int{2, 6, 4},
			listB: []int{1, 5},
			skipA: 3, skipB: 2,
		},
		{
			// 两条链表是同一条：第 0 步就相遇
			name:  "same_list",
			listA: []int{1, 2, 3},
			listB: []int{1, 2, 3},
			skipA: 0, skipB: 0,
		},
		{
			// B 是 A 的后缀：B 没有私有前缀
			name:  "b_is_suffix_of_a",
			listA: []int{1, 2, 3, 4},
			listB: []int{3, 4},
			skipA: 2, skipB: 0,
		},
		{
			name:  "a_is_suffix_of_b",
			listA: []int{7},
			listB: []int{5, 6, 7},
			skipA: 0, skipB: 2,
		},
		{
			// 私有前缀等长：第一圈就相遇，不需要换到另一条链表
			name:  "equal_prefixes",
			listA: []int{1, 2, 9, 9},
			listB: []int{3, 4, 9, 9},
			skipA: 2, skipB: 2,
		},
		{
			// 只在最后一个结点相交
			name:  "intersect_at_last_node",
			listA: []int{1, 2, 3, 4, 5},
			listB: []int{6, 5},
			skipA: 4, skipB: 1,
		},
		{
			// 不相交且长度相同：两个指针在第一圈末尾同时到达 nil
			name:  "equal_length_no_intersection",
			listA: []int{1, 2, 3},
			listB: []int{1, 2, 3},
			skipA: 3, skipB: 3,
		},
		{
			name:  "single_nodes_no_intersection",
			listA: []int{1},
			listB: []int{1},
			skipA: 1, skipB: 1,
		},
		{
			// 长度取上限 3*10^4，B 只有一个结点，即 A 的最后一个结点
			name:  "max_length_long_vs_short",
			listA: seq(30_000),
			listB: []int{30_000},
			skipA: 29_999, skipB: 0,
		},
		{
			// 长度都取上限且不相交：两个指针在第一圈末尾同时到达 nil
			name:  "max_length_equal_no_intersection",
			listA: seq(30_000),
			listB: seq(30_000),
			skipA: 30_000, skipB: 30_000,
		},
		{
			// 长度相差 1 且不相交：循环要走满 m + n + 1 步才在 nil 处相遇
			name:  "max_length_unequal_no_intersection",
			listA: seq(30_000),
			listB: seq(29_999),
			skipA: 30_000, skipB: 29_999,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			headA, headB, want, nodesA, nodesB := buildIntersecting(test.listA, test.listB, test.skipA, test.skipB)
			got := getIntersectionNode(headA, headB)
			if got != want {
				t.Errorf("getIntersectionNode(%s) = %p, want %p", test.name, got, want)
			}
			// 题目要求函数返回后链表保持原始结构
			if !sameStructure(headA, nodesA) || !sameStructure(headB, nodesB) {
				t.Errorf("getIntersectionNode(%s) modified the lists", test.name)
			}
		})
	}
}
