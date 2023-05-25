package leetcode

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

/*
思路：slow fast 快慢指针的思想
题目有点问题，思路要明确
*/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
