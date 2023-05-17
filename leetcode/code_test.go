package leetcode

import "testing"

func TestLeetCode(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	//l2 := ListNode{
	//	Val: 3,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  5,
	//			Next: nil,
	//		},
	//	},
	//}

	head := reverseList(&l1)
	for {
		if head == nil {
			break
		}
		println(head.Val)
		head = head.Next
	}
}
