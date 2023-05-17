package leetcode

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var root, nextNode *ListNode
	l1Nums := make([]int, 0)
	l2Nums := make([]int, 0)
	resNums := make([]int, 0)
	for {
		if l1 == nil {
			break
		}
		l1Nums = append(l1Nums, l1.Val)
		l1 = l1.Next
	}
	for {
		if l2 == nil {
			break
		}
		l2Nums = append(l2Nums, l2.Val)
		l2 = l2.Next
	}
	tmp := 0

	for {
		if len(l1Nums) == 0 && len(l2Nums) == 0 && tmp == 0 {
			break
		}
		n1, n2 := 0, 0

		if len(l1Nums) != 0 {
			n1 = l1Nums[len(l1Nums)-1]
			l1Nums = l1Nums[:len(l1Nums)-1]
		}
		if len(l2Nums) != 0 {
			n2 = l2Nums[len(l2Nums)-1]
			l2Nums = l2Nums[:len(l2Nums)-1]
		}

		sum := n1 + n2 + tmp
		if sum >= 10 {
			tmp = 1
			sum = sum % 10
		} else {
			tmp = 0
		}
		resNums = append(resNums, sum)

	}

	for i := len(resNums) - 1; i >= 0; i-- {
		tmpNode := &ListNode{
			Val:  resNums[i],
			Next: nil,
		}

		if root == nil {
			root = tmpNode
			nextNode = root
		} else {
			nextNode.Next = tmpNode
			nextNode = tmpNode
		}
	}

	return root
}
