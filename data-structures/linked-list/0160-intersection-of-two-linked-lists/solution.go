package solution

// ListNode 是 LeetCode 给出的单链表结点定义。
type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 返回两条单链表相交的起始结点，不相交时返回 nil。
// 两个指针各自走完自己的链表后，改从另一条链表的表头继续走；
// 这样两者走过的总长度相同，会在相交结点（或同时到达 nil）处相遇。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
