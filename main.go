package main

import (
	"fmt"
	"github.com/ilhamtubagus/leetcode-patterns/prefix_sum"
	"github.com/ilhamtubagus/leetcode-patterns/two_pointer"
)

// https://blog.algomaster.io/p/15-leetcode-patterns

func main() {
	_ = prefix_sum.SumElem([]int{1, 2, 3}, 1, 2)
	twoPointerSum := two_pointer.TwoNumberSum([]int{1, 2, 3, 4, 6}, 6)
	fmt.Println(twoPointerSum)
}
