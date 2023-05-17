package leetcode

//  Definition for singly-linked list.

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := new(ListNode)
	curNode := root //当前节点
	moreInt := 0    //进位值
	for l1 != nil || l2 != nil {
		curNode.Next = new(ListNode)
		curNode = curNode.Next
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum = sum + moreInt
		moreInt = sum / 10
		sum = sum % 10
		curNode.Val = sum
	}
	//处理最后一次进位情况
	if moreInt > 0 {
		curNode.Next = new(ListNode)
		curNode.Next.Val = moreInt
	}
	return root.Next
}
