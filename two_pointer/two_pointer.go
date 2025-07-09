package two_pointer

import (
	"sort"
)

// this only works if nums is sorted

func TwoNumberSum(nums []int, target int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	left, right := 0, len(nums)-1
	for i := 0; i < len(nums); {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		}

		if sum < target {
			// move left pointer to the right to get bigger value
			left++
		}
		if sum > target {
			right--
		}
	}

	return []int{left, right}
}
