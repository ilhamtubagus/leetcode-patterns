package next_greater_element

import "fmt"

type MonotonicStack []int

func (s *MonotonicStack) Pop() int {
	if len(*s) == 0 {
		return -1
	}
	lastElement := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return lastElement
}

func (s *MonotonicStack) Push(val int) {
	for len(*s) > 0 && val > (*s)[len(*s)-1] {
		s.Pop()
	}
	*s = append(*s, val)
}

func NextGreaterElement(nums1, nums2 []int) []int {
	// Map to store the next greater element for each number in nums2
	nextGreater := make(map[int]int)

	// Initialize our monotonic stack
	var stack MonotonicStack

	for _, num := range nums2 {
		// While the current number is greater than the top of the stack
		for len(stack) > 0 && num > stack[len(stack)-1] {
			smaller := stack.Pop()
			nextGreater[smaller] = num
			fmt.Println(nextGreater[smaller])
		}
		// Push current number to stack
		stack.Push(num)
	}

	// For remaining items in the stack, they have no next greater element
	for _, remaining := range stack {
		nextGreater[remaining] = -1
	}

	// Build the result for nums1 using the map
	result := make([]int, len(nums1))
	for i, num := range nums1 {
		result[i] = nextGreater[num]
	}

	return result
}
