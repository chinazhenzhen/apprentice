package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	v1Head := head
	var v1Before *ListNode
	v1Before = nil
	v2Head := head

	for i := 0; i < n; i++ {
		v2Head = v2Head.Next
	}

	for {
		if v2Head == nil {
			if v1Head != nil && v1Before != nil {
				v1Before.Next = v1Head.Next
			}
			if v1Before == nil {
				head = v1Head.Next
			}
			break
		}
		v1Before = v1Head
		v1Head = v1Head.Next
		v2Head = v2Head.Next
	}

	return head
}
