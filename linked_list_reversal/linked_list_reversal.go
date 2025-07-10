package linked_list_reversal

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	current := head
	var b strings.Builder

	for current != nil {
		if current.Next != nil {
			b.WriteString(fmt.Sprintf(" %d ->", current.Val))
		} else {
			b.WriteString(fmt.Sprintf(" %d", current.Val))
		}

		current = current.Next
	}

	return b.String()
}

func ReverseList(head *ListNode) *ListNode {
	current := head
	var previous *ListNode = nil

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
	}

	head = previous

	return head
}
