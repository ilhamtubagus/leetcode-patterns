package main

import (
	"fmt"
	"github.com/ilhamtubagus/leetcode-patterns/fast_and_slow_pointer"
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
	isCycled := fast_and_slow_pointer.HasCycle(&one)
	fmt.Println(isCycled)
}
