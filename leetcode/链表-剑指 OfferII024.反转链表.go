package leetcode

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var pre *ListNode

	for {
		if head == nil {
			break
		}
		next := head.Next
		head.Next = pre
		pre = head

		head = next
	}
	return pre
}
