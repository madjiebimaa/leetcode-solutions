package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var l ListNode
	for l1.Next != nil && l2.Next != nil {
		if l1.Val <= l2.Val {
			l.Next = l1
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
	}

	return &l
}
