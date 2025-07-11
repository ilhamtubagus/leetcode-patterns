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

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right || left <= 0 {
		return head
	}

	var previousLeft, rightNode, next *ListNode = nil, nil, nil
	dummy := &ListNode{Val: 0, Next: head}
	current := dummy

	if left == 1 {
		for i := 0; i < right; i++ {
			next = current.Next
			if i < left {
				previousLeft = current.Next
			}
			rightNode = current.Next
			current = next
		}
	} else {
		for i := 0; i < right; i++ {
			next = current.Next
			if i < left-1 {
				previousLeft = current.Next
			}
			rightNode = current.Next
			current = next
		}
	}

	// Detach
	if rightNode == nil {
		return head
	}
	var leftNode *ListNode = nil
	if left == 1 {
		leftNode = previousLeft
	} else {
		leftNode = previousLeft.Next
	}

	afterRight := rightNode.Next
	rightNode.Next = nil

	// Reverse
	current = leftNode
	var previous *ListNode = nil
	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	headReversed := previous

	// Attach
	if left == 1 {
		previousLeft = headReversed
	} else {
		previousLeft.Next = headReversed
	}

	current = previousLeft
	for current != nil {
		next = current.Next
		if current.Next == nil {
			current.Next = afterRight
		}
		current = next
	}

	if left == 1 {
		return previousLeft
	} else {
		return head
	}
}
