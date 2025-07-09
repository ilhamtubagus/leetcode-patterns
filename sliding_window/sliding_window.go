package sliding_window

func MaxSumOfSubArray(nums []int, k int) int {
	windowSum := 0
	// Calculate sum of the first window
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}

	maxSum := windowSum

	// Loop for the rest of the array, after k items
	for i := k; i < len(nums); i++ {
		// Slide the window by subtracting i-k item
		windowSum = windowSum - nums[i-k] + nums[i]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

/*
	MaxSumOfSubArray([]int{1, 2, 3, 4, 6, 1}, 3)

	1+2+3=6
	i = k = 3
	windowSum = 6 - nums[3-3] + nums [3]
	6 - 1 + 4 = 9

	i = 4
	windowSum = 9 - nums[4-3] + nums [4]
	9 - 2 + 6 = 13 ---> maxSum

	i = 5
	windowSum = 13 - nums[5-3] + nums[5]
	13 - 3 + 1 = 11
*/
