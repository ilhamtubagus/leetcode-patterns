package main

import (
	"fmt"
	"github.com/ilhamtubagus/leetcode-patterns/fast_and_slow_pointer"
	"github.com/ilhamtubagus/leetcode-patterns/linked_list_reversal"
	"github.com/ilhamtubagus/leetcode-patterns/next_greater_element"
	"github.com/ilhamtubagus/leetcode-patterns/prefix_sum"
	"github.com/ilhamtubagus/leetcode-patterns/sliding_window"
	"github.com/ilhamtubagus/leetcode-patterns/two_pointer"
)

// https://blog.algomaster.io/p/15-leetcode-patterns

func main() {
	_ = prefix_sum.SumElem([]int{1, 2, 3}, 1, 2)
	_ = two_pointer.TwoNumberSum([]int{1, 2, 3, 4, 6}, 6)
	_ = sliding_window.MaxSumOfSubArray([]int{1, 2, 3, 4, 6, 1}, 3)

	four := fast_and_slow_pointer.ListNode{Val: 4}
	three := fast_and_slow_pointer.ListNode{Val: 3, Next: &four}
	two := fast_and_slow_pointer.ListNode{Val: 2, Next: &three}
	//four.Next = &two
	one := fast_and_slow_pointer.ListNode{Val: 1, Next: &two}
	// 1 -> 2 -> 3 -> 4 -> 2
	_ = fast_and_slow_pointer.HasCycle(&one)

	// 5 -> 6 -> 7
	seven := linked_list_reversal.ListNode{Val: 7}
	six := linked_list_reversal.ListNode{Val: 6, Next: &seven}
	five := linked_list_reversal.ListNode{Val: 5, Next: &six}
	_ = linked_list_reversal.ReverseList(&five)

	// 8 -> 9 -> 10 -> 11 -> 12
	twelve := linked_list_reversal.ListNode{Val: 12}                // 5
	eleven := linked_list_reversal.ListNode{Val: 11, Next: &twelve} // 4
	ten := linked_list_reversal.ListNode{Val: 10, Next: &eleven}    // 3
	nine := linked_list_reversal.ListNode{Val: 9, Next: &ten}       // 2
	eight := linked_list_reversal.ListNode{Val: 8, Next: &nine}     // 1
	_ = linked_list_reversal.ReverseBetween(&eight, 1, 2)

	//sourceArr := []int{1, 7, 9, 5}
	//monotonicStack := next_greater_element.MonotonicStack{}
	//for _, num := range sourceArr {
	//	monotonicStack.Push(num)
	//}

	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Printf("%v\n", next_greater_element.NextGreaterElement(nums1, nums2))
}
