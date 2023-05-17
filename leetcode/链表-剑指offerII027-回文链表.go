package leetcode

//type ListNode struct {
//	Val int
//	Next *ListNode
//}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	nums := make([]int, 0)
	for {
		if head == nil {
			break
		}

		nums = append(nums, head.Val)
		head = head.Next
	}

	l := len(nums)
	for i := 0; i < l; i++ {
		if i < l-i-1 {
			if nums[i] != nums[l-i-1] {
				return false
			}
		}
	}
	return true
}
