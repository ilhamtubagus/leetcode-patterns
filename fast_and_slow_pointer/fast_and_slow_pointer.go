package fast_and_slow_pointer

// Definition for singly-linked list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	slow := head.Next
	fast := head.Next.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false
}

/*
1 -> 2 -> 3 -> 4 -> 2
*/
